package orders

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/XanderMoroz/BookStore/internal/middlewares"
)

// Регистрируем маршруты приложения
// Передаем указатель на сервер и указатель на базу данных
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Создаем указатель на структуру обработчика
	h := &handler{
		DB: db,
	}

	protected := r.Group("/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/orders", h.AddOrder)
	protected.GET("/orders", h.GetMyOrders)
	protected.GET("/orders/:id", h.GetMyOrder)
	protected.DELETE("/orders/:id", h.DeleteOrder)

}
