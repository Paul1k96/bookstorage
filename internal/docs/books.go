package docs

import "github.com/Paul1k96/bookstorage/internal/modules/books/controller"

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route POST /api/1/books/get_books books booksRequest
// Получение книг по автору.
// responses:
//	200: booksResponse

// swagger:parameters booksRequest
type booksRequest struct {
	// in:body
	// example: "Aubree Bayer"
	Body controller.BooksRequest
}

// swagger:response booksResponse
type booksResponse struct {
	// in:body
	Body controller.BooksResponse
}

// swagger:route POST /api/1/books/get_authors books authorsRequest
// Получение авторов по книге.
// responses:
//	200: authorsResponse

// swagger:parameters authorsRequest
type authorsRequest struct {
	// in:body
	// example: "Be am"
	Body controller.AuthorsRequest
}

// swagger:response authorsResponse
type authorsResponse struct {
	// in:body
	Body controller.AuthorsResponse
}
