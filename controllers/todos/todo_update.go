package todos

import (
	"net/http"

	"github.com/kurokuman/go-react-ToDoApp/databases"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var todo databases.Todo
	todoId, err := GetTodoId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error todo_id is should be a number")
		return
	}
	todo.Id = todoId

	message := todo.BindJson(c)
	if message != "" {
		c.JSON(http.StatusBadRequest, message)
		return
	}

	err = todo.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when trying to update")
	}

	c.JSON(http.StatusOK, todo)

}
