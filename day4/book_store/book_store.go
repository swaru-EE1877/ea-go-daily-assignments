package bookstore

import (
	"ea-go-daily-assignments/day4/book_store/models"
	"errors"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{Id: 1, Title: "Python", Price: 200},
	{Id: 2, Title: "Go", Price: 300},
	{Id: 3, Title: "Javascript", Price: 250},
	{Id: 4, Title: "Java", Price: 200},
}

func getBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, books)
}

func _findBook(books []models.Book, id int) (models.Book, error) {
	for _, book := range books {
		if book.Id == int(id) {
			return book, nil
		}
	}
	return models.Book{}, errors.New("Book not found")
}

func getBookDetails(ctx *gin.Context) {
	id, parsingErr := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if parsingErr != nil {
		ctx.String(http.StatusInternalServerError, "Book id is invalid")
		return
	}
	result, err := _findBook(books, int(id))
	if err != nil {
		ctx.String(http.StatusNotFound, "Book is not available")
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func addBook(ctx *gin.Context) {
	var newBook models.Book

	err := ctx.BindJSON(&newBook)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Invalid JSON")
	}

	books = append(books, newBook)
	ctx.JSON(http.StatusOK, books)
}

func _removeBookFromList(books []models.Book, id int) error {
	for index, book := range books {
		if book.Id == int(id) {
			books = books[0:index]
			if index+1 < len(books) {
				books = append(books, books[index+1:]...)
			}
			return nil
		}
	}
	return errors.New("Book not found")
}

func deleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Book id is invalid")
		return
	}
	err = _removeBookFromList(books, int(id))
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Successfully removed book from store")
}
