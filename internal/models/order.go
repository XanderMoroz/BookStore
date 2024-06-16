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
	ID        uuid.UUID `gorm:"type:char(255);primary_key" json:"ID"` // Уникальный идентификатор
	User      User      `json:"user"`
	UserID    uuid.UUID `gorm:"type:char(255);not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	Items     []Item `gorm:"foreignKey:OrderID"`
}

type OrderResponse struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Структура позиции в заказе
type Item struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(255);primary_key" json:"ID"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	// Заказ к которому прикреплена позиция
	Order   Order     `json:"order"`
	OrderID uuid.UUID `gorm:"type:char(255);not null" json:"order_id"`
	// Указатель на Книгу которая включена в позицию
	Book   Book      `json:"book"`
	BookID uuid.UUID `gorm:"type:char(255);not null" json:"book_id"`
}

type CreateItem struct {
	BookID   string `json:"book_id" validate:"required"`
	Quantity uint   `json:"quantity" validate:"required"`
	OrderID  string `json:"order_id" validate:"required"`
}
