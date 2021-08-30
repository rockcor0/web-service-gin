package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID		string	`json:"id"`
	Tittle	string	`json:"tittle"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []album{
	{ID: "1", Tittle: "Blue Train", Artist: "John Contrane", Price: 56.99},
	{ID: "2", Tittle: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Tittle: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumById locates the album whose ID value matches de id
// parametrer sent by the client, then returns that album as a response
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	//Recorrer la lista de albumes
	for _, a := range albums {
		if(a.ID == id){
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}