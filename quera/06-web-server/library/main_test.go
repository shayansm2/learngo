package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerCreation(t *testing.T) {
	s := getServer()
	assert.NotNil(t, s)
}

// post endpoint
func TestAddBook(t *testing.T) {
	title := "alice in wonderland"
	author := "JC"
	msg, err := addBook(title, author, false)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("added book %s by %s", strings.ToLower(title), strings.ToLower(author)), msg)
}

func TestInvalidAddBook1(t *testing.T) {
	_, err := addBook("", "", false)
	assert.NotNil(t, err)
	assert.Equal(t, "title or author cannot be empty", err.Error())
}

func TestInvalidAddBook2(t *testing.T) {
	_, err := addBook("something", "", false)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")
}

func TestAddInvalid3(t *testing.T) {
	_, err := addBook(" ", " ", false)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")

	_, err = addBook("  ", " no one  ", false)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")
}

func TestAddInvalid4(t *testing.T) {
	_, err := dummyAddBook("alaki ", " no one  ", false)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")
}

func TestAddDuplicateBooks1Incorrect(t *testing.T) {
	_, err := addBook("alice in wonderland", "jc", true)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book is already in the library")
}

func TestAddDuplicateBooks2(t *testing.T) {
	_, err := addBook("Alice In Wonderland", "jC", true)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book is already in the library")
}

// get endpoint
func TestGetBook(t *testing.T) {
	book, err := getBook("alice in wonderland", "JC")
	assert.Nil(t, err)
	assert.IsType(t, &testBook{}, book)
	assert.Equal(t, "alice in wonderland", book.Title)
	assert.Equal(t, "jc", book.Author)
}

func TestNotFoundGetBook(t *testing.T) {
	book, err := getBook("harry poter", "idk")
	assert.Nil(t, book)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book does not exist")
}

func TestInvalidGetBook(t *testing.T) {
	_, err := getBook("", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")
}

func TestBorrowedGetBook(t *testing.T) {
	reserveBook("alice in wonderland", "JC", true)
	defer reserveBook("alice in wonderland", "JC", false)
	_, err := getBook("alice in wonderland", "JC")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book is borrowed")
}

// put endpoint
func TestBorrowBook(t *testing.T) {
	msg, err := reserveBook("alice in wonderland", "JC", true)
	defer reserveBook("alice in wonderland", "JC", false)
	assert.Nil(t, err)
	assert.Equal(t, msg, "you have borrowed this Book successfully")
}

func TestReturnBook(t *testing.T) {
	reserveBook("alice in wonderland", "JC", true)
	msg, err := reserveBook("alice in wonderland", "JC", false)
	assert.Nil(t, err)
	assert.Equal(t, msg, "thank you for returning this Book")
}

func TestInvalidBorrowBook(t *testing.T) {
	_, err := reserveBook("", "", true)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")
}

func TestNonReturnBook(t *testing.T) {
	_, err := reserveBook("alice in wonderland", "JC", false)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book is already in the library")
}

func TestBorrowedBorrowBook(t *testing.T) {
	reserveBook("alice in wonderland", "JC", true)
	defer reserveBook("alice in wonderland", "JC", false)
	_, err := reserveBook("alice in wonderland", "JC", true)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book is already borrowed")
}

func TestNoBorrowFlag(t *testing.T) {
	_, err := reserveBook("alice in wonderland", "JC")
	assert.NotNil(t, err)
	assert.Equal(t, "borrow value cannot be empty", err.Error())
}

// del endpoint
func TestDeleteBook(t *testing.T) {
	msg, err := delBook("alice in wonderland", "JC")
	assert.Nil(t, err)
	assert.Equal(t, msg, "successfully deleted")
}

func TestInvaluidDeleteBook(t *testing.T) {
	_, err := delBook("", "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "title or author cannot be empty")
}

func TestNotFoundDeleteBook(t *testing.T) {
	_, err := delBook("1984", "someone")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "this Book does not exist")
}

func TestAddDuplicateBooks(t *testing.T) {
	// Add book once
	_, err := addBook("test book", "test author", false)
	assert.Nil(t, err)

	// Try to add the same book again
	msg, err := addBook("test book", "test author", false)
	assert.Nil(t, err)
	assert.Equal(t, "this Book is already in the library", msg)
}