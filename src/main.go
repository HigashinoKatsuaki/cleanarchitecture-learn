package main

import (
	"cleanarchitecture-learn/src/injector"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := sqlx.Connect("mysql", "root:root@(localhost:9000)/development?parseTime=true&loc=Asia%2FTokyo")

	if err != nil {
		log.Fatal(err)
	}

	todoHandler := injector.InjectUsecase(db)

	// ルーティング
	e.GET("/todos", todoHandler.All())

	// サーバー起動
	e.Start(":9001")
}
