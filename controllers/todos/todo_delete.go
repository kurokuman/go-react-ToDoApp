package todos

import (
	"fmt"
	"net/http"

	"github.com/kurokuman/go-react-ToDoApp/databases"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	todoId, err := GetTodoId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error todo_id is should be a number")
		return
	}
	todo := databases.Todo{Id: todoId}

	err = todo.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error when trying to delete todo(id=%d)", todo.Id))
		return
	}

	c.JSON(http.StatusOK, "success delete")
}
