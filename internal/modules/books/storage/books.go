package storage

import (
	"github.com/Paul1k96/bookstorage/internal/models"
	"github.com/jmoiron/sqlx"
)

type BookStorage struct {
	db *sqlx.DB
}

func NewBookStorage(db *sqlx.DB) BookStorager {
	return &BookStorage{db: db}
}

func (bs *BookStorage) GetBooksByAuthor(author string) ([]models.Book, error) {
	var books []models.Book

	stmt := `
		SELECT b.book_id, b.book_name
		FROM authors_books ab
		INNER JOIN books b ON b.book_id = ab.book_id
		INNER JOIN authors a ON a.author_id = ab.author_id
		WHERE a.author_name = ?
	`

	if err := bs.db.Select(&books, stmt, author); err != nil {
		return nil, err
	}

	return books, nil
}

func (bs *BookStorage) GetAuthorsByBook(book string) ([]models.Author, error) {
	var authors []models.Author

	stmt := `
		SELECT a.author_id, a.author_name
		FROM authors_books ab
		INNER JOIN books b ON b.book_id = ab.book_id
		INNER JOIN authors a ON a.author_id = ab.author_id
		WHERE b.book_name = ?
	`

	if err := bs.db.Select(&authors, stmt, book); err != nil {
		return nil, err
	}

	return authors, nil

}
