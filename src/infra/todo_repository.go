package infra

import (
	"cleanarchitecture-learn/src/domain/model"
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
