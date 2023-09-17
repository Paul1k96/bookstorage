package modules

import (
	bs "github.com/Paul1k96/bookstorage/internal/modules/books/storage"
	"github.com/jmoiron/sqlx"
)

type Storages struct {
	book bs.BookStorager
}

// NewStorages сборка всех хранилищ в одной структуре для инициализации
func NewStorages(sql *sqlx.DB) *Storages {
	return &Storages{
		book: bs.NewBookStorage(sql),
	}
}
