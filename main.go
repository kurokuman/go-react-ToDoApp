package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kurokuman/bookstore_users-api/logger"
)

func main() {
	router := gin.Default()

	logger.Info("about to start the application")
	// router.POST("/todos/", todos.Create)
	// router.GET("/todos/:todo_id", todos.Get)
	// router.PUT("/todos/:todo_id", todos.Update)
	// router.DELETE("/todos/:todo_id", todos.Delete)
	router.Run(":8080")
}
