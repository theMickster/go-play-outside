package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{Id: "41cbc53b-9434-4321-8542-6a1fdac1fc8f", Title: "Up All Night", Artist: "Kip Moore", Price: 9.99},
	{Id: "dfea4654-d5da-425a-9866-ab6c9dd10d7e", Title: "Wild Ones", Artist: "Kip Moore", Price: 11.99},
	{Id: "c8b3a3e1-bec4-4a8a-9879-2ae50b187cbc", Title: "Slowheart", Artist: "Kip Moore", Price: 7.99},
	{Id: "be9fa363-e9c8-4a75-b2c8-68c19a457710", Title: "Wild World", Artist: "Kip Moore", Price: 10.39},
	{Id: "69e12637-1b77-465e-b0ba-d2b70831f92e", Title: "Damn Love", Artist: "Kip Moore", Price: 11.19},
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORS())
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8081")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}
}
