package service

import (
	"github.com/Paul1k96/bookstorage/internal/infrastructure/errors"
	"github.com/Paul1k96/bookstorage/internal/modules/books/storage"
	"log"
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
		log.Println("GetBooksByAuthor err", err)
		return GetBooksByAuthorOut{
			ErrorCode: errors.BookServiceGetBooksByAuthorErr,
		}
	}

	return GetBooksByAuthorOut{Book: books}

}

func (bs *BookService) GetAuthorsByBook(in GetAuthorsByBookIn) GetAuthorsByBookOut {
	authors, err := bs.storage.GetAuthorsByBook(in.Name)
	if err != nil {
		log.Println("GetBooksByAuthor err", err)
		return GetAuthorsByBookOut{
			ErrorCode: errors.BookServiceGetAuthorsByBookErr,
		}
	}

	return GetAuthorsByBookOut{Author: authors}
}
