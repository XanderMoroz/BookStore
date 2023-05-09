package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"BookStore/internal/models"
)

// Инициализация базы данных
func Init() *gorm.DB {
	err := godotenv.Load("./envs/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := "POSTGRES"
	DbUser := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("%s://%s:%s@%s:%s/%s", DbUser, DbUser, DbPassword, DbHost, DbPort, DbName)

	log.Println("Осуществляем подключение к базе")
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		fmt.Println("Невозможно подключиться к базе ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Подключение к базе прошло успешно ", Dbdriver)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
