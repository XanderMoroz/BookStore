package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/XanderMoroz/BookStore/internal/middlewares"
	"github.com/XanderMoroz/BookStore/internal/models"
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
	routes := r.Group("/users")

	//routes.POST("/", h.AddBook)
	routes.POST("/register", h.Register) // Создание нового пользователя
	routes.POST("/login", h.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", models.CurrentUser)

}
