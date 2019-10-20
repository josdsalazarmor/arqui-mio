package main

import (
	"github.com/gin-gonic/gin"
	"github.com/multimedia_ms/infrastructure/api/handler"
	"github.com/multimedia_ms/infrastructure/persistance"
)

func main() {
	conn := persistance.NewSQLDB()
	defer conn.Close()
	entryRepo := persistance.NewEntryRepository(conn)

	eH := handler.NewEntryHandler(entryRepo)

	r := gin.Default()

	entriesGroup := r.Group("/files")
	entriesGroup.GET("", eH.GetEntries)
	entriesGroup.POST("", eH.CreateEntry)
	entriesGroup.GET(":userId", eH.GetEntry)
	entriesGroup.PUT(":userId", eH.UpdateEntry)
	entriesGroup.DELETE(":userId", eH.DeleteEntry)

	/*
	tagsGroup := r.Group("/tags")
	tagsGroup.GET("", tH.GetTags)
	tagsGroup.POST("", tH.CreateTag)
	tagsGroup.GET(":id", tH.GetTag)
	tagsGroup.PUT(":id", tH.UpdateTag)
	tagsGroup.DELETE(":id", tH.DeleteTag)
	*/

	r.Run(":8081")
}
