package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album
func postAlbums(c *gin.Context){
	var newAlbum album

	//Llamar BindJson para mapear el JSON recibido a album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Agregar el  nuevo álbum
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumById get an album
/*func getAlbumById(c *gin.Context){
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}*/

// deleteAlbums delete an album
/*func deleteAlbums(c *gin.Context){
	oldAlbum := c.Param("id")

	for _, album := range albums {
		if album.ID == oldAlbum {
			albums = albums
		}
	}
}*/

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}