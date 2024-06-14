package utils

import (
	"log"

	"github.com/google/uuid"

	"github.com/XanderMoroz/BookStore/db"
	"github.com/XanderMoroz/BookStore/internal/models"
)

// Извлекаем пользователя по ID
func GetOrderByIDFromDB(orderUUID string) models.Order {

	db := db.DB
	var order models.Order

	// Извлекаем заказ по ID
	result := db.First(&order, "ID = ?", orderUUID)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve order: " + result.Error.Error())
	}

	if order.ID == uuid.Nil {
		// handle error
		panic("failed to retrieve order: " + result.Error.Error())
	}

	log.Println("Заказ — успешно извлечен:")
	log.Printf("	ID: <%s>\n", order.ID)
	log.Printf("	Имя Заказчика: <%s>\n", order.User.Name)

	return order
}
