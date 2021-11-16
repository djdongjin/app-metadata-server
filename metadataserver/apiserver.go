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

// PersistMetadata receives yaml-formatted string, converts it to a
// Metadata object, and persists the object using persister.
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

	if err = server.DataStore.Persist(metadata); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusCreated, "Metadata persisted.")
}

// RetrieveMetadata receives a query string, parse it to subqueries, and returns
// all Metadata selected by persister that matches subqueries.
func (server *ApiServer) RetrieveMetadata(c *gin.Context) {
	reqBody, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	subqueries, err := common.ParseQuery(string(reqBody))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	allMetadata, err := server.DataStore.Retrieve(subqueries)
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

// GetMetadata receives a title from GET request url, get the Metadata from
// persister, and returns it to client.
func (server *ApiServer) GetMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Get(title); ok {
		c.String(http.StatusOK, common.MetadataToYamlString(metadata))
	} else {
		c.String(http.StatusNotFound, "Metadata not found.")
	}
}

// DeleteMetadata receives a title from GET request url, let persister
// delete that metadata if exists.
func (server *ApiServer) DeleteMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Delete(title); ok {
		c.String(http.StatusOK, "Medadata deleted: \n\n"+common.MetadataToYamlString(metadata))
	} else {
		c.String(http.StatusNotFound, "Metadata not found.")
	}
}
