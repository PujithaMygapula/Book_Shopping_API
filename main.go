package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type book struct {
	BookID      int    `json:"bookID"`
	BookName    string `json:"bookName"`
	Author      string `json:"author"`
	ReviewScore int    `json:"reviewScore"`
	SoldCount   int    `json:"soldCount"`
}

var books = []book{
	{BookID: 1, BookName: "C", Author: "Shivani", ReviewScore: 5, SoldCount: 10},
	{BookID: 2, BookName: "Python", Author: "Pujitha", ReviewScore: 8, SoldCount: 16},
	{BookID: 3, BookName: "Java", Author: "Sailaja", ReviewScore: 4, SoldCount: 12},
	{BookID: 4, BookName: "Php", Author: "Manju", ReviewScore: 7, SoldCount: 34},
	{BookID: 5, BookName: "Go", Author: "Mounika", ReviewScore: 9, SoldCount: 25},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func addBook(context *gin.Context) {
	var newBook book

	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getBook_ID(id int) (*book, error) {
	for i, b := range books {
		if b.BookID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found with given book id")
}

func getBookById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	book, err := getBook_ID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found with given Book Id"})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func getBook_Name(name string) (*[]book, error) {
	var sameBooks = []book{}
	for i := 0; i < len(books); i++ {
		if strings.Contains(strings.ToLower(books[i].BookName), strings.ToLower(name)) {
			sameBooks = append(sameBooks, books[i])
		}
	}
	if len(sameBooks) != 0 {
		return &sameBooks, nil
	}
	return nil, errors.New("no book data found with the provided book name")
}

func getBookByName(context *gin.Context) {
	bookName := context.Param("name")
	sameBooks, err := getBook_Name(bookName)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Book Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, sameBooks)
}

func updateBookInfo(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	book, err := getBook_ID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	book.BookName = "JavaScript"
	book.ReviewScore = 10
	book.SoldCount = 34

	context.IndentedJSON(http.StatusOK, book)
}

func deleteByBookId(id int) (string, error) {
	msg := false
	for i := 0; i < len(books); i++ {
		if books[i].BookID == id {
			books[i] = books[len(books)-1]
			books = books[:len(books)-1]
			msg = true
		}
	}
	if msg {
		return "Book Record Deleted Successfully", nil
	} else {
		return " ", errors.New("book Not found with the given bookID")
	}
}

func deleteBookById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	book, err := deleteByBookId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found with provided Book ID"})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/getbooks", getBooks)
	router.GET("/getbookbyid/:id", getBookById)
	router.GET("/getbookbyname/:name", getBookByName)
	router.POST("/addbook", addBook)
	router.PATCH("updatebookbyid/:id", updateBookInfo)
	router.DELETE("/deletebookbyid/:id", deleteBookById)
	router.Run("localhost:1211")
}
