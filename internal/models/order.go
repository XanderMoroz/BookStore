package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order represents the model for an order
// Default table name will be `orders`
type Order struct {
	// gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"ID"` // Уникальный идентификатор
	User      User      `json:"author"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	Items     []Item `gorm:"foreignKey:OrderID"`
}

// Структура позиции в заказе
type Item struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primary_key" json:"ID"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	// Заказ к которому прикреплена позиция
	Order   Order  `json:"order"`
	OrderID uint64 `gorm:"type:uint64;not null" json:"order_id"`
	// Указатель на Книгу которая включена в позицию
	Book   Book      `json:"book"`
	BookID uuid.UUID `gorm:"type:char(36);not null" json:"book_id"`
}
