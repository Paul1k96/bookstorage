package router

import (
	"net/http"

	"github.com/Paul1k96/bookstorage/internal/modules"
	"github.com/go-chi/chi/v5"
)

func NewApiRouter(controllers *modules.Controllers) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/1", func(r chi.Router) {
			r.Route("/books", func(r chi.Router) {
				bookController := controllers.Books
				r.Post("/get_books", bookController.GetBooks)
				r.Post("/get_authors", bookController.GetAuthors)
			})
		})
	})

	return r
}
