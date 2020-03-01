package repository

import "cleanarchitecture-learn/src/domain/model"

type TodoRepository interface {
	All() []*model.Todo
	Save(*model.Todo) error
}
