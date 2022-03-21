package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

//	backtick içindeki json bölümlerinin amacı struct ı serialize etmek (jsona çevirmek django da da serialize olarak geçiyor)
//	Fieldların büyük harfle başlama nedenide go da export edilecek değerin büyük harfle başlaması
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Apology of Socrates", Author: "Plato", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "Brave New World", Author: "Aldous Huxley", Quantity: 6},
}

// c *gin.Context kısaca gelen request
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func bookById(c *gin.Context) {
	id := c.Param("id") //path i books/id değeri yapan kod
	book, err := getBookById(id)

	if err != nil {
		//gin.H kendi jsonumuzu yazmamızı sağlar
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
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

	return nil, errors.New("book not found!")
}

func createBook(c *gin.Context) {
	var newBook book
	// newBookun pointerından değerlerini request ettiğimizde error veriyorsa düz return yapıyoruz
	//c.BindJSON response otomatik olarak yolluyor
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.Run("localhost:8080")

}
