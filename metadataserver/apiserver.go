package metadataserver

import (
	"net/http"

	"github.com/djdongjin/app-metadata-server/persisters"
	"github.com/gin-gonic/gin"
)

// ApiServer is the main struct handling HTTP requests. It holds a DataStore that
// implements the Persister interface.
type ApiServer struct {
	DataStore persisters.Persister
}

func (server *ApiServer) PersistMetadata(c *gin.Context) {
	reqBody, err := c.GetRawData()
	if err != nil {
		return
	}

	metadata, err := YamlStringToMetadata(string(reqBody))
	_ = server.DataStore.Persist(metadata)

	c.String(http.StatusCreated, "Metadata persisted.")
}

func (server *ApiServer) RetrieveMetadata(c *gin.Context) {
	queryBody, err := c.GetRawData()
	if err != nil {
		return
	}

	allMetadata := server.DataStore.Retrieve(string(queryBody))
	var res string
	for _, md := range allMetadata {
		res += MetadataToYamlString(md) + "\n\n"
	}

	c.String(http.StatusOK, res)
}

func (server *ApiServer) GetMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Get(title); ok {
		c.String(http.StatusOK, MetadataToYamlString(metadata))
	} else {
		c.String(http.StatusNotFound, "Metadata not found.")
	}
}

func (server *ApiServer) DeleteMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Delete(title); ok {
		c.String(http.StatusOK, "Medadata deleted: \n\n"+MetadataToYamlString(metadata))
	} else {
		c.String(http.StatusNotFound, "Metadata not found.")
	}
}
