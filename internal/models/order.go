package models

import "time"

// Order represents the model for an order
// Default table name will be `orders`
type Order struct {
	// gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

type Item struct {
	// gorm.Model
	BookID      uint   `gorm:"foreignkey:BookID" json:"bookId"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `gorm:"foreignkey:OrderID" json:"-"`
}
