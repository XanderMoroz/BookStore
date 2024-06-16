package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model            // adds ID, created_at etc.
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"ID"` // Уникальный идентификатор
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PublishedAt int       `json:"published_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
	// Категории к которым принадлежит статья
	Genres []*Genre `gorm:"many2many:book_genres;"`
}

type BookResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublishedAt int    `json:"published_at"`
}

// Определяем структуру тела запроса на создание экземпляра книги
type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublishedAt int    `json:"published_at"`
}
