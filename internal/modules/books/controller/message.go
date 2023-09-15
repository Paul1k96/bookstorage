package controller

import "github.com/Paul1k96/bookstorage/internal/models"

//go:generate easytags $GOFILE
type BooksRequest struct {
	Name string `json:"name"`
}

type BooksResponse struct {
	Success   bool          `json:"success"`
	ErrorCode int           `json:"error_code"`
	Message   string        `json:"message"`
	Books     []models.Book `json:"books"`
}

type AuthorsRequest struct {
	Name string `json:"name"`
}

type AuthorsResponse struct {
	Success   bool            `json:"success"`
	ErrorCode int             `json:"error_code"`
	Message   string          `json:"message"`
	Authors   []models.Author `json:"authors"`
}
