package modules

import (
	"github.com/Paul1k96/bookstorage/internal/infrastructure/responder"
	bs "github.com/Paul1k96/bookstorage/internal/modules/books/controller"
)

type Controllers struct {
	Books bs.Booker
}

func NewControllers(services *Services, r responder.Responder) *Controllers {
	return &Controllers{
		Books: bs.NewBooks(services.Book, r),
	}
}
