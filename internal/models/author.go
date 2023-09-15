package models

//go:generate easytags $GOFILE json,db
type Author struct {
	ID   int    `json:"id" db:"author_id"`
	Name string `json:"name" db:"name"`
}
