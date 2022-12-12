package main

import (
	"fmt"
	"main/src/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Welcome to the server!")

	e := echo.New()

	e.GET("/", handler.Yallo)

	e.GET("/cats/:data", handler.GetCats)

	e.Start(":8000")
}
