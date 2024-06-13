package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/XanderMoroz/BookStore/internal/middlewares"
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

	routes.POST("/register", h.Register) // Создание нового пользователя
	routes.POST("/login", h.Login)
	routes.GET("/logout", h.Logout)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", h.CurrentUser)

}
