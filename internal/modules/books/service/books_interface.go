package service

import "github.com/Paul1k96/bookstorage/internal/models"

type BookServicer interface {
	GetBooksByAuthor(in GetBooksByAuthorIn) GetBooksByAuthorOut
	GetAuthorsByBook(in GetAuthorsByBookIn) GetAuthorsByBookOut
}

type GetBooksByAuthorIn struct {
	Name string
}

type GetBooksByAuthorOut struct {
	Book      []models.Book
	ErrorCode int
}

type GetAuthorsByBookIn struct {
	Name string
}

type GetAuthorsByBookOut struct {
	Author    []models.Author
	ErrorCode int
}
