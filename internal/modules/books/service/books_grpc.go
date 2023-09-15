package service

import (
	"context"

	"github.com/Paul1k96/bookstorage/pkg/grpc/books"
)

type BooksServiceGRPC struct {
	books BookServicer
	books.UnimplementedBookServiceGRPCServer
}

func NewBooksServiceGRPC(books BookServicer) *BooksServiceGRPC {
	return &BooksServiceGRPC{books: books}
}

func (b *BooksServiceGRPC) GetBooksByAuthor(ctx context.Context, in *books.GetBooksByAuthorIn) (*books.GetBooksByAuthorOut, error) {
	elem := GetBooksByAuthorIn{
		Name: in.Author,
	}

	res := b.books.GetBooksByAuthor(elem)

	out := make([]*books.Book, 0, len(res.Book))

	for i := range res.Book {
		out = append(out, &books.Book{
			Id:   int32(res.Book[i].ID),
			Name: res.Book[i].Name,
		})
	}

	return &books.GetBooksByAuthorOut{
		Books:     out,
		ErrorCode: int32(res.ErrorCode),
	}, nil
}

func (b *BooksServiceGRPC) GetAuthorsByBook(ctx context.Context, in *books.GetAuthorsByBookIn) (*books.GetAuthorsByBookOut, error) {
	elem := GetAuthorsByBookIn{
		Name: in.Book,
	}

	res := b.books.GetAuthorsByBook(elem)

	out := make([]*books.Author, 0, len(res.Author))

	for i := range res.Author {
		out = append(out, &books.Author{
			Id:   int32(res.Author[i].ID),
			Name: res.Author[i].Name,
		})
	}

	return &books.GetAuthorsByBookOut{
		Authors:   out,
		ErrorCode: int32(res.ErrorCode),
	}, nil

}
