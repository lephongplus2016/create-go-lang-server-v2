package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the webside")
}

// lấy query param
func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	// param
	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s & %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{"cat": catName, "type": catType})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to let's us know if you want to json or string data",
	})
}
