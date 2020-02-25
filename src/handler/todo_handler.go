package handler

import (
	"cleanarchitecture-learn/src/usecase"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type TodoHandler struct {
	UC usecase.TodoUsecase
}

func (h *TodoHandler) All() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoList := h.UC.All()

		buf, err := json.Marshal(todoList)

		if err != nil {
			log.Fatal(err)
		}

		return c.String(http.StatusOK, string(buf))
	}
}
