package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/XanderMoroz/BookStore/config"
	"github.com/XanderMoroz/BookStore/internal/models"
)

// Указатель на БД
// (он будет осуществлять запросы)
var DB *gorm.DB

// Инициализация базы данных
func Init() *gorm.DB {

	env := config.NewEnv()

	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)

	log.Printf("Подключаемся к БД: <%s> ...", env.Dbdriver)
	log.Printf("По адресу: <%s> ...", DSN)

	var err error
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Printf("Error connecting to database : error=%v", err)
		return nil
	} else {
		log.Println("... успешно")
		log.Printf("	DB_HOST: <%s>", env.DBHost)
		log.Printf("	DB_PORT: <%s>", env.DBPort)
		log.Printf("	DB_NAME: <%s>", env.DBName)
		log.Printf("	DB_USER: <%s>", env.DBUser)
	}

	DB = db

	log.Printf("Устанавливаем миграции в БД...")
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Order{}, &models.Item{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	log.Printf("... успешно!")

	return db
}
