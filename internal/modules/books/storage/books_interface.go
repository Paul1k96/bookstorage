package storage

import "github.com/Paul1k96/bookstorage/internal/models"

type BookStorager interface {
	GetBooksByAuthor(author string) ([]models.Book, error)
	GetAuthorsByBook(book string) ([]models.Author, error)
}
