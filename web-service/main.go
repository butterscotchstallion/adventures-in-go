package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 50},
	{ID: "2", Title: "Kind of Blue", Artist: "Miles Davis", Price: 25},
	{ID: "3", Title: "Hoshizora Feeling", Artist: "Charisma.com", Price: 50},
}

func getAlbums(c *gin.Context) {
	c.JSON(200, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"alive": true})
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/healthcheck", healthCheckHandler)
	return router
}

func main() {
	router := setUpRouter()
	router.Run("localhost:8080")
}
