package todos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurokuman/bookstore_users-api/logger"
	"github.com/kurokuman/go-react-ToDoApp/databases"
)

func GetTodoId(c *gin.Context) (int64, error) {
	todoIdParam := c.Param("todo_id")

	todoId, err := strconv.ParseInt(todoIdParam, 10, 64)
	if err != nil {
		logger.Error("error when todo_id convert to int64 from string", err)
		return 0, err
	}

	return todoId, nil
}

func BindJson(c *gin.Context, todo databases.Todo) (*databases.Todo, string) {
	var message string
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		logger.Error("error to should bind json", err)

		c.JSON(http.StatusBadRequest, "invalid json body")
		message = "invalid json body"
		return &todo, message
	}

	message = todo.Validate()
	if message != "" {
		logger.Info(message)
		return &todo, message
	}

	return &todo, message
}
