package todos

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kurokuman/bookstore_users-api/logger"
	"github.com/kurokuman/go-react-ToDoApp/databases"
)

func Create(c *gin.Context) {
	logger.Info("access Create")

	var todo databases.Todo

	todo.Title = "test title"
	todo.Content = "test content"

	todo.Create()

	fmt.Println(todo)
}
