package service

import (
	"github.com/Paul1k96/bookstorage/internal/infrastructure/errors"
	"github.com/Paul1k96/bookstorage/internal/modules/books/storage"
)

type BookService struct {
	storage storage.BookStorager
}

func NewBookService(storage storage.BookStorager) BookServicer {
	return &BookService{storage: storage}
}

func (bs *BookService) GetBooksByAuthor(in GetBooksByAuthorIn) GetBooksByAuthorOut {
	books, err := bs.storage.GetBooksByAuthor(in.Name)
	if err != nil {
		return GetBooksByAuthorOut{
			ErrorCode: errors.BookServiceGetBooksByAuthorErr,
		}
	}

	return GetBooksByAuthorOut{Book: books}

}

func (bs *BookService) GetAuthorsByBook(in GetAuthorsByBookIn) GetAuthorsByBookOut {
	authors, err := bs.storage.GetAuthorsByBook(in.Name)
	if err != nil {
		return GetAuthorsByBookOut{
			ErrorCode: errors.BookServiceGetAuthorsByBookErr,
		}
	}

	return GetAuthorsByBookOut{Author: authors}
}
