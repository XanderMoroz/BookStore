package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// Создаем новый экземпляр книги
	var book models.Book

	// Пытаемся найти экземпляры
	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
