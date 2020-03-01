package infra

import (
	"cleanarchitecture-learn/src/domain/model"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	DB *sqlx.DB
}

func (r *TodoRepository) All() []*model.Todo {
	var todoList []*model.Todo
	rows, err := r.DB.Queryx("SELECT * FROM todos")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var todo model.Todo

		err := rows.StructScan(&todo)

		if err != nil {
			log.Fatal(err)
		}

		todoList = append(todoList, &todo)
	}

	return todoList
}

func (r *TodoRepository) Save(todo *model.Todo) error {
	insert := "INSERT INTO todos(status, task) VALUES(?, ?)"

	if todo.Status == "" {
		todo.Status = "in progress"
	}

	fmt.Println(todo)

	_, err := r.DB.Exec(insert, todo.Status, todo.Task)

	return err
}
