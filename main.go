package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search Of Lost Item", Author: "Vaishnav Singh", Quantity: 2},
	{ID: "2", Title: "In Search Of Found Item", Author: "Bhote Singh", Quantity: 5},
	{ID: "3", Title: "In Search Of Get Lost Item", Author: "Vishu Singh", Quantity: 6},
}

func getBooks(c *gin.Context) { //c *gin.Context this reiterates and store  values according to specific request

	//Indented Json is just giving formatted json
	c.IndentedJSON(http.StatusOK, books)

}

// the .bind json method is what will handel sending the json response
func createbook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {

		return

	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}
func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createbook)

	router.Run("localhost:8080")
}
