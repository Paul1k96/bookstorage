package modules

import bs "github.com/Paul1k96/bookstorage/internal/modules/books/controller"

type Controllers struct {
	Books bs.Booker
}

func NewControllers(services *Services) *Controllers {
	return &Controllers{
		Books: bs.NewBooks(services.Book),
	}
}
