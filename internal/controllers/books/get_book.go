package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// @Summary		get an book by ID
// @Description Get an book by ID
// @Tags 		Books
// @ID			get-book-by-id
// @Produce		json
// @Param		id				path		string					true	"Book ID"
// @Success		200				{object}	models.BookResponse
// @Failure		404				{object}	[]string
// @Router		/books/{id} 	[get]
func (h handler) GetBook(c *gin.Context) {

	// Извлекаем из контекста значение параметра ID
	id := c.Param("id")

	// Инициализируем пустой экземпляр книги
	var book models.Book

	// Пытаемся найти экземпляр книги с полученным значением ID
	if result := h.DB.First(&book, id); result.Error != nil {
		// c.AbortWithError(http.StatusNotFound, result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Book Not Found",
		})
		return
	}

	// Отправляем в контекст экземпляр книги
	c.JSON(http.StatusOK, &book)
}
