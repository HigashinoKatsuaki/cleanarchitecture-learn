package injector

import (
	"cleanarchitecture-learn/src/domain/repository"
	"cleanarchitecture-learn/src/handler"
	"cleanarchitecture-learn/src/infra"
	"cleanarchitecture-learn/src/usecase"

	"github.com/jmoiron/sqlx"
)

func InjectDB(db *sqlx.DB) repository.TodoRepository {
	return &infra.TodoRepository{DB: db}
}

func InjectRepository(db *sqlx.DB) usecase.TodoUsecase {
	return usecase.TodoUsecase{Repo: InjectDB(db)}
}

func InjectUsecase(db *sqlx.DB) handler.TodoHandler {
	return handler.TodoHandler{UC: InjectRepository(db)}
}
