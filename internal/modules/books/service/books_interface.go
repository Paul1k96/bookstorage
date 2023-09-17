package service

import "github.com/Paul1k96/bookstorage/internal/models"

//go:generate mockgen -source books_interface.go -destination=mocks/book_mock.go

// BookServicer интерфейс описывает сервис слой сервиса books
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
