package main

import (
	"BookStore/db"
	"BookStore/internal/controllers/books"
	"BookStore/internal/controllers/users"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	// Инициализируем базу данных на основе GORM;
	h := db.Init()

	// Регистрируем маршруты приложений
	books.RegisterRoutes(r, h)
	users.RegisterRoutes(r, h)

	// Извлекаем переменную окружения PORT
	port := os.Getenv("PORT")

	// Запускаем сервер на указанном порту
	r.Run(port)
}
