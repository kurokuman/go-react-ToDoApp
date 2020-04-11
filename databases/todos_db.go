package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kurokuman/bookstore_users-api/logger"
)

var (
	db *sql.DB
)

const (
	queryInsertTodo  = "insert into todos(title, contents) values(?,?) "
	queryGetAllTodos = "select * from todos"
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
