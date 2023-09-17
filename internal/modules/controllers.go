package modules

import (
	"github.com/Paul1k96/bookstorage/internal/infrastructure/responder"
	bs "github.com/Paul1k96/bookstorage/internal/modules/books/controller"
)

type Controllers struct {
	Books bs.Booker
}

// NewControllers сборка всех контроллеров в одну структуру для инициализации
func NewControllers(services *Services, r responder.Responder) *Controllers {
	return &Controllers{
		Books: bs.NewBooks(services.Book, r),
	}
}
