package utils

import (
	"log"

	"github.com/google/uuid"

	"github.com/XanderMoroz/BookStore/db"
	"github.com/XanderMoroz/BookStore/internal/models"
)

// Извлекаем пользователя по ID
func GetUserByIDFromDB(userUUID string) models.User {

	db := db.DB
	var user models.User

	// Извлекаем статью вместе с автором и категориями
	result := db.First(&user, "ID = ?", userUUID)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve user: " + result.Error.Error())
	}

	if user.ID == uuid.Nil {
		// handle error
		panic("failed to retrieve user: " + result.Error.Error())
	}

	log.Println("Пользователь — успешно извлечен:")
	log.Printf("	ID: <%s>\n", user.ID)
	log.Printf("	Имя: <%s>\n", user.Name)
	log.Printf("	Username: <%s>\n", user.Username)

	return user
}
