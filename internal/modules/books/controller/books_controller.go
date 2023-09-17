package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Paul1k96/bookstorage/internal/infrastructure/errors"
	"github.com/Paul1k96/bookstorage/internal/infrastructure/responder"
	"github.com/Paul1k96/bookstorage/internal/modules/books/service"
)

// Booker интерфейс описывает контроллер сервиса books
type Booker interface {
	GetBooks(http.ResponseWriter, *http.Request)
	GetAuthors(http.ResponseWriter, *http.Request)
}

type Books struct {
	service service.BookServicer
	responder.Responder
}

func NewBooks(service service.BookServicer, r responder.Responder) Booker {
	return &Books{service: service, Responder: r}
}

// GetBooks метод для получения книг по автору. Приём и отправка данных происходит в формате JSON
func (b *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	var req BooksRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	out := b.service.GetBooksByAuthor(service.GetBooksByAuthorIn{Name: req.Name})

	if out.ErrorCode != errors.NoError {
		b.OutputJSON(w, BooksResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Message:   "Get books error",
		})
		return
	}

	b.OutputJSON(w, BooksResponse{
		Success: true,
		Message: "success",
		Books:   out.Book,
	})
}

// GetAuthors метод для получения авторов по книге. Приём и отправка данных происходит в формате JSON
func (b *Books) GetAuthors(w http.ResponseWriter, r *http.Request) {
	var req AuthorsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		b.ErrorBadRequest(w, err)
		return
	}

	out := b.service.GetAuthorsByBook(service.GetAuthorsByBookIn{Name: req.Name})

	if out.ErrorCode != errors.NoError {
		b.OutputJSON(w, BooksResponse{
			Success:   false,
			ErrorCode: out.ErrorCode,
			Message:   "Get authors error",
		})
		return
	}

	b.OutputJSON(w, AuthorsResponse{
		Success: true,
		Message: "success",
		Authors: out.Author,
	})
}
