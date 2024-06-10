package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// Определяем функцию обработчик для получения списка
func (h handler) GetBooks(c *gin.Context) {

	// Создаем пустой срез экземпляров
	var books []models.Book

	// Пытаемся найти экземпляры
	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Отправляем в контекст список экземпляров
	c.JSON(http.StatusOK, &books)
}
