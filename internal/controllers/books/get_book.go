package books

import (
	"BookStore/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Определяем функцию обработчик для получения книги по ID
func (h handler) GetBook(c *gin.Context) {
	// Извлекаем из контекста значение параметра ID
	id := c.Param("id")

	// Инициализируем пустой экземпляр книги
	var book models.Book

	// Пытаемся найти экземпляр книги с полученным значением ID
	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// Отправляем в контекст экземпляр книги
	c.JSON(http.StatusOK, &book)
}
