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

	// use native go
	e.POST("/cats", handler.AddCat)
	// use native go
	e.POST("/dogs", handler.AddDog)

	// use three party
	e.POST("/hamsters", handler.AddHamsters)

	e.Start(":8000")
}
