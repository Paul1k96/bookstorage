package router

import (
	"net/http"

	"github.com/Paul1k96/bookstorage/internal/modules"
	"github.com/Paul1k96/bookstorage/internal/router"
	"github.com/go-chi/chi/v5"
)

func NewRouter(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()
	setSwagger(r)

	r.Mount("/", router.NewApiRouter(controllers))
	return r
}

func setSwagger(r *chi.Mux) {
	r.Get("/swagger", swaggerUI)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
	})
}
