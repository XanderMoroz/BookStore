package utils

import (
	"log"

	"github.com/XanderMoroz/BookStore/db"
	"github.com/XanderMoroz/BookStore/internal/models"
	"github.com/google/uuid"
)

// Извлекаем все Жанры и авторов
func GetAllGenresFromDB() []models.Genre {

	db := db.DB
	var genres []models.Genre // genre slice

	result := db.Preload("Books").Find(&genres)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve genres: " + result.Error.Error())
	}

	log.Println("Список жанров — успешно извлечен:")
	for _, genre := range genres {
		log.Printf("Genre ID: <%d>, Title: <%s>\n", genre.ID, genre.Title)
	}
	return genres
}

// Извлекаем жанра по названию
func GetGenreByNameFromDB(name string) models.Genre {

	db := db.DB
	var genre models.Genre // genre slice

	result := db.Preload("Books").First(&genre, "Title = ?", name)

	if result.Error != nil {
		// handle error
		panic("failed to retrieve genre: " + result.Error.Error())
	}

	if genre.ID == uuid.Nil {
		// handle error
		panic("failed to retrieve genre: " + result.Error.Error())
	}

	log.Println("Жанр — успешно извлечен:")
	log.Printf("Genre ID: <%d>, Title: <%s>\n", genre.ID, genre.Title)

	return genre
}
