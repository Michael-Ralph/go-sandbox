package main

import (
	"fmt"
	"time"
)

type Book struct {
	ID         int
	Title      string
	Author     string
	ISBN       string
	Available  bool
	BorrowedBy string
	BorrowedAt time.Time
}

// String method for Book (value reciever)
// This implements the Stringer interface
func (b Book) String() string {
	status := "Available"
	if !b.Available {
		status = fmt.Sprintf("Borrowed by %s on %s", b.BorrowedBy, b.BorrowedAt.Format("2006-01-02"))
	}
	return fmt.Sprintf("ID: %d | %s by %s | status: %s", b.ID, b.Title, b.Author, status)
}

// Borrow method (pointer reciever) - modified the struct
func (b *Book) Borrow(BorrowerName string) error {

}

// Return method (pointer reciever) - modified the struct
func (b *Book) Return() error {

}

// IsOverdue method (value reciever) - reads but doesn't modify
func (b Book) IsOverdue() bool {

}

type Library struct {
	Name   string
	Books  []Book
	nextID int
}

// NewLibrary constructor function
func NewLibrary(name string) *Library {

}

// AddBook method (pointer reciever) - adds a new book to the library
func (l *Library) AddBook(title, author, isbn string) {

}

// FindBookByID method (pointer receiver) - returns pointer to allow modification
func (l *Library) FindBookByID(id int) *Book {

}

// SearchBooks method (value receiver) - searches by title or author
func (l Library) SearchBooks(query string) []Book {

}

// ListAllBooks method (value receiver)
func (l Library) ListAllBooks() {

}

// GetOverdueBooks method (value receiver)
func (l Library) GetOverdueBooks() []Book {

}

func main() {

}
