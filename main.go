package main

import (
	"log"

	"github.com/djdongjin/app-metadata-server/metadataserver"
	"github.com/djdongjin/app-metadata-server/persisters"
	"github.com/gin-gonic/gin"
)

func main() {
	dataStore, err := persisters.New("InMemoryPersister")
	if err != nil {
		log.Fatalf("Server initialization failed: %s\n", err.Error())
	}

	server := metadataserver.ApiServer{DataStore: dataStore}
	router := gin.Default()

	router.POST("/persist", server.PersistMetadata)
	router.POST("/retrieve", server.RetrieveMetadata)
	router.GET("/get/:title", server.GetMetadata)
	router.GET("/delete/:title", server.DeleteMetadata)

	router.Run("localhost:8080")
}
