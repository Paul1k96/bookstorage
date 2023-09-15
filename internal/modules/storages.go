package modules

import (
	bs "github.com/Paul1k96/bookstorage/internal/modules/books/storage"
	"github.com/jmoiron/sqlx"
)

type Storages struct {
	book bs.BookStorager
}

func NewStorages(sql *sqlx.DB) *Storages {
	return &Storages{
		book: bs.NewBookStorage(sql),
	}
}
