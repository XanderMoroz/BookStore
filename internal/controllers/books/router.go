package books

import (
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
	routes := r.Group("/books")

	routes.POST("/", h.AddBook)         // Создание новой книги
	routes.GET("/", h.GetBooks)         // Извлечение списка всех книг
	routes.GET("/:id", h.GetBook)       // Извлечение книги по ID
	routes.PUT("/:id", h.UpdateBook)    // Изменение полей книги по ID
	routes.DELETE("/:id", h.DeleteBook) // Удаление книги по ID

}
