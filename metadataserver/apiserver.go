package metadataserver

import (
	"fmt"

	// "github.com/djdongjin/app-metadata-server/common"
	"github.com/djdongjin/app-metadata-server/persisters"
	"github.com/gin-gonic/gin"
)

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
}

func (server *ApiServer) RetrieveMetadata(c *gin.Context) {
	queryBody, err := c.GetRawData()
	if err != nil {
		return
	}

	fmt.Println(string(queryBody))
}

func (server *ApiServer) GetMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Get(title); ok {
		fmt.Println(MetadataToYamlString(metadata))
	} else {
		fmt.Println("no Metadata found")
	}
}

func (server *ApiServer) DeleteMetadata(c *gin.Context) {
	title := c.Param("title")

	if metadata, ok := server.DataStore.Get(title); ok {
		fmt.Println(MetadataToYamlString(metadata))
	} else {
		fmt.Println("no Metadata found")
	}
}
