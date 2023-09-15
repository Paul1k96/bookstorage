package models

//go:generate easytags $GOFILE json,db
type Book struct {
	ID   int    `json:"id" db:"book_id"`
	Name string `json:"name" db:"name"`
}
