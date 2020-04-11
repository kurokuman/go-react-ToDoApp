package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurokuman/bookstore_users-api/logger"
	"github.com/kurokuman/go-react-ToDoApp/databases"
)

func Create(c *gin.Context) {
	logger.Info("access Create")

	var todo databases.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil {
		logger.Error("error to should bind json", err)
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}

	message := todo.Validate()
	if message != "" {
		logger.Info(message)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	err = todo.Create()
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, todo)
}
