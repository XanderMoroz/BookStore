package utils

import (
	"log"

	"github.com/XanderMoroz/BookStore/db"
	"github.com/XanderMoroz/BookStore/internal/models"
	"github.com/google/uuid"
	// "github.com/google/uuid"
)

// Извлекаем статью по ID
func GetBookByIDFromDB(bookID string) models.Book {

	db := db.DB
	var book models.Book

	// Извлекаем статью вместе с автором и категориями
	result := db.Preload("Genres").First(&book, "ID = ?", bookID)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve book: " + result.Error.Error())
	}

	if book.ID == uuid.Nil {
		// handle error
		panic("failed to retrieve book: " + result.Error.Error())
	} else {
		// handle success
		log.Println("Cтатья — успешно извлечена:")
		log.Printf("	ID: <%d>\n", book.ID)
		log.Printf("	Название: <%s>\n", book.Title)
		log.Printf("	Описание: <%s>\n", book.Description)
		log.Printf("	Автор: <%s>\n", book.Author)
		log.Printf("	Год издания: <%v>\n", book.PublishedAt)
	}

	return book
}
