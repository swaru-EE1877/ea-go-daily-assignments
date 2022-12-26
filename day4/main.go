package main

import bookstore "ea-go-daily-assignments/day4/book_store"

func main() {
	router := bookstore.BookStore()
	router.Run(":8080")
}
