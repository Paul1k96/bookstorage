package storage

import "github.com/Paul1k96/bookstorage/internal/models"

//go:generate mockgen -source books_interface.go -destination=mocks/book_mock.go

type BookStorager interface {
	GetBooksByAuthor(author string) ([]models.Book, error)
	GetAuthorsByBook(book string) ([]models.Author, error)
}
