package databases

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kurokuman/bookstore_users-api/logger"
)

var (
	db *sql.DB
)

const (
	queryInsertTodo  = "insert into todos(title, contents) values(?,?) "
	queryUpdateTodo  = "update todos set title=?, contents=? where id=?"
	queryGetAllTodos = "select * from todos"
	queryGetTodo     = "select * from todos where id=?"
	queryDeleteTodo  = "delete from todos where id=?"
)

func init() {
	var err error

	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/react_todos")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("database successfully configured")
	}
}

func GetAll() ([]Todo, error) {
	var todos []Todo
	rows, err := db.Query(queryGetAllTodos)
	if err != nil {
		logger.Error("error when trying to query to get all ", err)
		return todos, err
	}
	defer rows.Close()

	var todo Todo
	for rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		if err != nil {
			logger.Error("error when trying to scan rows", err)
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (todo *Todo) Get() error {
	stmt, err := db.Prepare(queryGetTodo)
	if err != nil {
		logger.Error("error when trying to prepare get todo statement", err)
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(todo.Id)
	err = result.Scan(&todo.Id, &todo.Title, &todo.Content)
	if err != nil {
		logger.Error("error when trying to get todo by id", err)
		return err
	}

	return nil

}

func (todo *Todo) Create() error {

	stmt, err := db.Prepare(queryInsertTodo)
	if err != nil {
		logger.Error("error when trying to prepare save todo statement ", err)
		return err
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(todo.Title, todo.Content)
	if err != nil {
		logger.Error("error when trying save todo", err)
		return err
	}

	todoId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("err when trying to get last todo id after creating a new todo", err)
		return err
	}

	todo.Id = todoId

	return nil
}

func (todo *Todo) Update() error {

	stmt, err := db.Prepare(queryUpdateTodo)
	if err != nil {
		logger.Error("error when trying to prepare update todo statement ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Content, todo.Id)
	if err != nil {
		logger.Error("error when trying update todo", err)
		return err
	}

	return nil
}

func (todo *Todo) BindJson(c *gin.Context) string {
	var message string
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		logger.Error("error to should bind json", err)
		message = "invalid json body"
		return message
	}

	message = todo.Validate()
	if message != "" {
		logger.Info(message)
		return message
	}

	return message
}

func (todo *Todo) Delete() error {
	stmt, err := db.Prepare(queryDeleteTodo)
	if err != nil {
		logger.Error("error when trying to prepare delete todo statement", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Id)
	if err != nil {
		logger.Error("error when trying to delete todo", err)
		return err
	}

	return nil

}
