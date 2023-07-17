package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create ablum struct to run $go run . in directory to establish server, new window $curl http://localhost:8080/albums
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 14.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 12.99},
	{ID: "3", Title: "AstroWorld", Artist: "Travis Scott", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// add get function for first endpoint
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//call BindJSON to bind recieved JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// get album by id
// parameter sent by client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//loop over list of ablums looking for album whose ID value matches the param
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
