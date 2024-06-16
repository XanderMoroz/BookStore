package genres

import (
	"github.com/XanderMoroz/BookStore/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Формируем структуру обработчика базы данных
type handler struct {
	DB *gorm.DB
}

// Регистрируем маршруты приложения
// Передаем указатель на сервер и указатель на базу данных
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Создаем указатель на структуру обработчика
	h := &handler{
		DB: db,
	}
	routes := r.Group("/genres")

	routes.GET("/", h.GetAllGenres) // Извлечение списка всех  жанров

	protected := r.Group("/genres")
	protected.Use(middlewares.JwtAuthMiddleware())
	// Создание нового жанра книги
	protected.POST("/", h.CreateNewGenre)
	protected.POST("/add_book", h.AddBookToGenre)
	protected.POST("/remove_book", h.RemoveBookFromGenre)
}
