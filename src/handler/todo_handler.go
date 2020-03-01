package handler

import (
	"cleanarchitecture-learn/src/domain/model"
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

func (h *TodoHandler) Save() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := new(model.Todo)

		err := c.Bind(t)

		if err != nil {
			log.Fatal(err)
		}

		err = h.UC.Save(t)

		if err != nil {
			log.Fatal(err)
		}

		return c.String(http.StatusOK, "")
	}
}
