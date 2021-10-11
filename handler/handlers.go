package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/kenyaachon/go-urlshortener/store"
	"github.com/kenyaachon/go-urlshortener/shortener"
)

//Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(context *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := context.ShouldBindJSON(&creationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808"
	context.JSON(200, gin.H{
		"message": "short url created succesfully",
		"short_url": host + shortUrl,
	})
}
func HandleShortUrlRedirect(context *gin.Context){
	shortUrl := context.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	context.Redirect(302, initialUrl)
}