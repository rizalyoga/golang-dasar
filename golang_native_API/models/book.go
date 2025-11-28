package models

import "time"

type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	AuthorID    uint      `json:"author_id"`
	Author      Author    `grom:"foreignKey:AuthorID" json:"author" `
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BooksResponse struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	AuthorID    uint               `json:"-"`
	Author      AuthorBookResponse `grom:"foreignKey:AuthorID" json:"author" `
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
