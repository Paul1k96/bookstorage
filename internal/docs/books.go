package docs

import "github.com/Paul1k96/bookstorage/internal/modules/books/controller"

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route GET /api/1/books/get_books books booksRequest
// Получение книг по автору.
// responses:
//
//	200: booksResponse

// swagger:parameters BooksRequest
type booksRequest struct {
	// in:body
	Body controller.BooksRequest
}

// swagger:response BooksResponse
type booksResponse struct {
	// in:body
	Body controller.BooksResponse
}

// swagger:route GET /api/1/books/get_authors books authorsRequest
// Получение авторов по книге.
// responses:
//	200: authorsResponse

// swagger:parameters AuthorsRequest
type authorsRequest struct {
	// in:body
	Body controller.AuthorsRequest
}

// swagger:response AuthorsResponse
type authorsResponse struct {
	// in:body
	Body controller.AuthorsResponse
}
