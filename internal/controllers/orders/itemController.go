package orders

import (
	"log"
	"net/http"

	"github.com/XanderMoroz/BookStore/internal/models"
	"github.com/XanderMoroz/BookStore/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Формируем структуру обработчика базы данных
type itemHandler struct {
	DB *gorm.DB
}

// @Summary        create new Book
// @Description    Creating Book in DB with given request body
// @Tags           Books
// @Accept         json
// @Produce        json
// @Param          request         			body        models.CreateItem    true    "Введите данные статьи"
// @Success        201              		{string}    string
// @Failure        400              		{string}    string    "Bad Request"
// @securityDefinitions.apiKey token
// @in				header
// @name			Authorization
// @Security		JWT
// @Router         /orders/add_book			[post]
func (h handler) AddBookToOrder(c *gin.Context) {

	// Определяем тело запроса
	body := models.CreateItem{}
	// Извлекаем JWT токен из куки пользователя
	user_id, err := utils.ExtractUserIDFromToken(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := utils.GetUserByIDFromDB(user_id)

	// Пытаемся получить тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := utils.GetOrderByIDFromDB(body.OrderID)

	// Если заказа нет возвращаем ошибку
	if order.ID == uuid.Nil {
		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "Order not found",
			"data":    nil,
		})
		return
	}

	// Если пользователь не является автором статьи
	if order.UserID != currentUser.ID {
		log.Println(order.UserID)
		log.Println(currentUser.ID)
		c.JSON(http.StatusNoContent, gin.H{
			"success": false,
			"message": "You can update your oun resourses only",
			"data":    nil,
		})
		return
	}

	book := utils.GetBookByIDFromDB(body.BookID)

	// Создаем новый экземпляр позиции в заказе
	var itemInOrder models.Item

	// Присваиваем ему значения из тела запроса и уникальный ID
	itemInOrder.ID = uuid.New()
	itemInOrder.Order = order
	itemInOrder.OrderID = order.ID
	itemInOrder.Book = book
	itemInOrder.BookID = book.ID
	itemInOrder.Quantity = body.Quantity

	log.Printf("Добавляем в заказ новую позицию:")
	log.Printf("	ID: <%+v>\n", itemInOrder.ID)
	log.Printf("	Заказ ID: <%s>\n", itemInOrder.OrderID)
	log.Printf("	Автор ID: <%+v>\n", itemInOrder.Order.UserID)
	log.Printf("	Заказываемый товар: <%v>\n", itemInOrder.Book)
	log.Printf("	Количество: <%v>\n", itemInOrder.Quantity)

	// Пытаемся создать экземпляр книги
	if result := h.DB.Create(&itemInOrder); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	// Отправляем в контекст сообщение об успешном создании экземпляра книги
	h.DB.Model(&order).Association("Items").Append(&itemInOrder)

	log.Println("Заказ — успешно обновлена:")
	log.Printf("	ID заказа: <%s>\n", itemInOrder.OrderID)
	log.Printf("	ID товара: <%s>\n", itemInOrder.BookID)
	log.Printf("	Количество: <%v>\n", itemInOrder.Quantity)

	// Return success message
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book added to order",
		"data":    &itemInOrder,
	})
	return
}
