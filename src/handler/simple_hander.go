package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/src/models"
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

func AddCat(c echo.Context) error {
	cat := models.Cat{}

	// nhớ đóng truy cập ở cuối function
	defer c.Request().Body.Close()

	// truy cập vào body request
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to reading body request: %s", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to reading body request: %s", err))
	}

	// đọc được rồi
	// giờ ép kiểu json sang struct gán cho cat
	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed to Unmarshal: %s", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to Unmarshal: %s", err))
	}

	// Đã lấy được body param và gán vào cat, giờ ta tùy ý xử lý
	log.Printf("This is your cat: %s", cat)

	return c.String(http.StatusOK, "we get your cat")
}

func AddDog(c echo.Context) error {
	dog := models.Dog{}

	// nhớ đóng truy cập ở cuối function
	defer c.Request().Body.Close()

	// Ta sử dụng cách 2: JSON Encode
	// ta copy body của c.Request() và lưu và &dog
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed to reading body request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Đã lấy được body param và gán vào dog, giờ ta tùy ý xử lý
	log.Printf("This is your dog: %s", dog)

	return c.String(http.StatusOK, "we get your dog")
}

// ít hiệu quả, chậm hơn 2 cách trên
func AddHamsters(c echo.Context) error {
	hamster := models.Hamster{}

	// cách lấy body json này sử dụng hàm do echo hỗ trợ
	// c.Bind() lấy đc tất cả, nhưng yêu cầu phải truyền reference
	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed to reading body request: %s", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to reading body request: %s", err))
	}

	log.Printf("This is your hamster: %s", hamster)

	return c.String(http.StatusOK, "we get your hamster")
}
