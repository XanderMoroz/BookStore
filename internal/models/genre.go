package models

import (
	"time"

	"gorm.io/gorm"
)

// Структура статьи
type Genre struct {
	gorm.Model
	ID    uint64 `gorm:"primary_key;auto_increment" json:"ID"` // Уникальный идентификатор
	Title string `gorm:"size:255;not null" json:"title"`       // Название категории
	// Список статей связанных с категорией
	Books []*Book `gorm:"many2many:book_genres;"`
}

// CreateCategoryBody
// @Description Тело запроса для создания жанра
type CreateGenreBody struct {
	Title string `json:"title" validate:"required"`
}

// GenreResponse
// @Description Тело ответа после извлечения жанра
type GenreResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

// CreateCategoryBody
// @Description Тело запроса для создания жанра
type AddBookToGenreBody struct {
	ArticleID    string `json:"article_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
}
