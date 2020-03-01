package usecase

import (
	"cleanarchitecture-learn/src/domain/model"
	"cleanarchitecture-learn/src/domain/repository"
)

type TodoUsecase struct {
	Repo repository.TodoRepository
}

func (u *TodoUsecase) All() []*model.Todo {
	return u.Repo.All()
}

func (u *TodoUsecase) Save(t *model.Todo) error {
	return u.Repo.Save(t)
}
