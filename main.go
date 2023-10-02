package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Articles)
}

func handleRequests() {
	log.Print("Launched")

	router := gin.Default()
	router.GET("/all", func(c *gin.Context) {
		// Handle the request and send a JSON response
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	router.GET("/health", func(c *gin.Context) {
		// Handle the request and send a JSON response
		c.JSON(http.StatusOK, gin.H{"message": "Online: Version 1.0"})
	})

	router.Run(":3600")
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
