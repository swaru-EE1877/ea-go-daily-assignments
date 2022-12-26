package bookstore_test

import (
	"bytes"
	bookstore "ea-go-daily-assignments/day4/book_store"
	"ea-go-daily-assignments/day4/book_store/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	var books []models.Book
	json.Unmarshal(data, &books)
	assert.Equal(t, books, []models.Book{
		{Id: 1, Title: "Python", Price: 200},
		{Id: 2, Title: "Go", Price: 300},
		{Id: 3, Title: "Javascript", Price: 250},
		{Id: 4, Title: "Java", Price: 200},
	})
}
func TestPostBooks(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	newBook := models.Book{Id: 6, Title: "Ruby", Price: 250}

	newBookJson, _ := json.Marshal(newBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(newBookJson))
	r.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	var books []models.Book
	json.Unmarshal(data, &books)
	assert.Equal(t, books, []models.Book{
		{Id: 1, Title: "Python", Price: 200},
		{Id: 2, Title: "Go", Price: 300},
		{Id: 3, Title: "Javascript", Price: 250},
		{Id: 4, Title: "Java", Price: 200},
		{Id: 6, Title: "Ruby", Price: 250},
	})
}
func TestShouldReturnErrorIfInvalidJsonPassedWhenPostBooks(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	newBook := "invalid book"

	newBookJson, _ := json.Marshal(newBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(newBookJson))
	r.ServeHTTP(rec, req)

	assert.Equal(t, 400, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	var books []models.Book
	json.Unmarshal(data, &books)
	assert.Equal(t, books, []models.Book(nil))
}
func TestGetBookDetail(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/3", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	var book models.Book
	json.Unmarshal(data, &book)

	assert.Equal(t, book, models.Book{Id: 3, Title: "Javascript", Price: 250})
}
func TestShouldReturnErrorIfInvalidIdPassedWhenGetBookDetail(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/sget", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 500, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "Book id is invalid")
}
func TestShouldReturnErrorIfBookNotFoundWhenGetBookDetail(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/14", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 404, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "Book is not available")
}
func TestDeleteBook(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/3", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "Successfully removed book from store")
}
func TestShouldReturnErrorIfBookNotFoundWhenDeleteBook(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/8", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 404, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "Book not found")
}
func TestShouldReturnErrorIfInvalidIdPassedWhenDeleteBook(t *testing.T) {
	r := bookstore.BookStore()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/erw", nil)
	r.ServeHTTP(rec, req)

	assert.Equal(t, 500, rec.Code)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "Book id is invalid")
}
