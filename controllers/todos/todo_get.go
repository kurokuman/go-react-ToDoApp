package todos

import (
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
