package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Структура жанра
type Genre struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:char(255);primary_key" json:"ID"`       // Уникальный идентификатор
	Title string    `gorm:"size:255;not null;uniqueIndex" json:"title"` // Название категории
	Books []*Book   `gorm:"many2many:book_genres;"`                     // Список статей связанных с категорией
}

// @Description Тело запроса для создания жанра
type CreateGenreBody struct {
	Title string `json:"title" validate:"required"`
}

// @Description Тело ответа после извлечения жанра
type GenreResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// @Description Тело запроса для связываня книги с жанром
type BookToGenreBody struct {
	BookID string `json:"book_id" validate:"required"`
	Genre  string `json:"genre" validate:"required"`
}
