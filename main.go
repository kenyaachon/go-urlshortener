package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kenyaachon/go-urlshortener/store"
	"github.com/kenyaachon/go-urlshortener/handler"

)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})


	router.POST("/create-short-url", func(context *gin.Context) {
		handler.CreateShortUrl(context)
	})

	router.GET("/:shortUrl", func(context *gin.Context) {
		handler.HandleShortUrlRedirect(context)
	})

	//Note that store initialization happens here
	store.InitializeStore()

	err := router.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}