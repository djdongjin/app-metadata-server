package metadataserver

import (
	"net/http"

	"github.com/djdongjin/app-metadata-server/common"
	"github.com/djdongjin/app-metadata-server/persisters"
	"github.com/gin-gonic/gin"
)

// ApiServer is the main struct handling HTTP requests. It holds a DataStore that
// implements the Persister interface.
type ApiServer struct {
	DataStore persisters.Persister
}

// PersistMetadata receives yaml-formatted raw string from POST request, converts
// it to a Metadata object, and persists the object in the underlying storage.
func (server *ApiServer) PersistMetadata(c *gin.Context) {
	reqBody, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	metadata, err := common.YamlStringToMetadata(string(reqBody))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err = metadata.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_ = server.DataStore.Persist(metadata)

	c.String(http.StatusCreated, "Metadata persisted.")
}

// RetrieveMetadata receives a query string from POST request, and returns
// filtered metadata to the client. Notice the query is passed to the Persister,
// where the actual retrieval happens.
func (server *ApiServer) RetrieveMetadata(c *gin.Context) {
	queryBody, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	allMetadata, err := server.DataStore.Retrieve(string(queryBody))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var res string
	for _, md := range allMetadata {
		res += common.MetadataToYamlString(md) + "\n\n"
	}

	c.String(http.StatusOK, res)
}

// GetMetadata receives a title from GET request url and returns that
// metadata to the client.
func (server *ApiServer) GetMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Get(title); ok {
		c.String(http.StatusOK, common.MetadataToYamlString(metadata))
	} else {
		c.String(http.StatusNotFound, "Metadata not found.")
	}
}

// DeleteMetadata receives a title from GET request and let the Persister
// delete that metadata, if exists.
func (server *ApiServer) DeleteMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Delete(title); ok {
		c.String(http.StatusOK, "Medadata deleted: \n\n"+common.MetadataToYamlString(metadata))
	} else {
		c.String(http.StatusNotFound, "Metadata not found.")
	}
}
