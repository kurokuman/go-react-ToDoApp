package todos

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurokuman/go-react-ToDoApp/databases"
)

func GetAll(c *gin.Context) {
	todos, err := databases.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when trying get all todos")
		return
	}

	c.JSON(http.StatusOK, todos)
}

func Get(c *gin.Context) {
	todoId, err := GetTodoId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error todo_id is should be a number")
		return
	}

	todo := databases.Todo{Id: todoId}

	err = todo.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error when trying get todo(todo_id is:%d)", todo.Id))
		return
	}

	c.JSON(http.StatusOK, todo)

}
