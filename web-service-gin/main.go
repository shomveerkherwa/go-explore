package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"strconv"
)

// creating a struct for type ALBUM
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// initializing the field
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// sets up an association in which getAlbums handles requests to the /albums endpoint path.
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:index", removeByIndex)

	router.Run("localhost:8080")
}

// handler to return all the albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// loop over the list and get the one that matches the id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// attempt to implement HTTP delete
func removeByIndex(c *gin.Context) {
	index := c.Param("index")
	position, error := strconv.Atoi(index)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	fmt.Println("before ", albums)
	albums = append(albums[:position], albums[position+1:]...)
	fmt.Println("after ", albums)
	c.IndentedJSON(http.StatusOK, albums)
}
