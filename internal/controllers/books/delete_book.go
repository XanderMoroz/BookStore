package books

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/XanderMoroz/BookStore/internal/models"
)

// @Summary		delete a book by ID
// @Description Delete a book by ID
// @ID			delete-book-by-id
// @Tags 		Books
// @Produce		json
// @Param		id				path		string		true	"Book ID"
// @Success		200				{object}	[]string
// @Failure		404				{object}	[]string
// @Router		/books/{id} 	[delete]
func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// Создаем новый экземпляр книги
	var book models.Book

	// Пытаемся найти экземпляры
	if result := h.DB.First(&book, id); result.Error != nil {
		// c.AbortWithError(http.StatusNotFound, result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   result.Error,
		})
		return
	}

	h.DB.Delete(&book)

	c.JSON(http.StatusNotFound, gin.H{
		"success": true,
		"message": "Book Deleted",
	})

}
