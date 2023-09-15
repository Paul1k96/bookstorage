package modules

import bs "github.com/Paul1k96/bookstorage/internal/modules/books/service"

type Services struct {
	Book bs.BookServicer
}

func NewServices(storages *Storages) *Services {
	return &Services{
		Book: bs.NewBookService(storages.book),
	}
}
