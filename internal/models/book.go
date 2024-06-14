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
	// Genries []*Genry `gorm:"many2many:book_genries;"`
}

// // Структура статьи
// type Article struct {
// 	gorm.Model
// 	ID        uint64    `gorm:"primary_key;auto_increment" json:"ID"` // Уникальный идентификатор
// 	User      User      `json:"author"`                               // Автор статьи
// 	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`  // Уникальный идентификатор автора статьи
// 	Title     string    `gorm:"size:255;not null" json:"title"`       // Название статьи
// 	Content   string    `gorm:"size:255;not null;" json:"content"`    // Содержание статьи
// 	CreatedAt time.Time `gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime"`
// 	DeletedAt gorm.DeletedAt
// 	// Категории к которым принадлежит статья
// 	Categories []*Category `gorm:"many2many:article_categories;"`
// 	Comments   []Comment   `gorm:"foreignKey:ArticleID"`
// }

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
