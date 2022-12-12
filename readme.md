# Creating Golang WebServer With Echo

## Response Json

-   trong 1 response api method luôn có 2 kiểu là

```Go
return c.String(http.StatusOK,...)
return c.JSON(http.StatusOK,...)
```

-   Với `JSON` có 3 cách tạo ra 1 json để trả về.

1. Ta thông qua 1 struct

```Go
// User
type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

// Handler
func(c echo.Context) error {
  u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.JSON(http.StatusOK, u)
}

```

2. Ta tạo ra Json bằng map tự tạo

```Go
return c.JSON(http.StatusOK, map[string]string{"cat": catName, "type": catType})
```

3. Ta tạo encode

Chưa rõ lắm cách này

## Defer

-   https://marcofranssen.nl/the-use-of-defer-in-go
-   `Defer` will always be triggered at the end of a function.
-   So even if the code panics in some location of the executing code it will guarantee the deferred code will be executed. A panic in Go is an unhandled error and causes the program execution to get halted. After a panic a defer will be executed. Panic is not recommended to use it for exception handling. It is better to handle exceptions using Golang error object.

```Go
package main

import "fmt"

func main() {
     deferExample()
     fmt.Println("Returned from deferExample.")
}

func deferExample() {
     defer fmt.Println("Deferred log.")
     fmt.Println("Print line.")
}
```

-   Kết quả:

```Bash
go run main.go
Print line.
Deferred log.
Returned from deferExample.
```

> Defer là 1 trigger luôn thực thi ở cuối function run time, kể cả có panic error xảy ra.

## Parsing JSON From Request

-   Cách 1: nhanh nhất

```Go
func AddCat(c echo.Context) error {
	cat := models.Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to reading body request: %s", err))
	}


	err = json.Unmarshal(b, &cat)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to Unmarshal: %s", err))
	}

	return c.String(http.StatusOK, "we get your cat")
}
```

-   Cách 2: nhanh

```Go
func AddDog(c echo.Context) error {
	dog := models.Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, "we get your dog")
}

```

-   Cách 3: ít hiệu quả, chậm hơn 2 cách trên

```Go
func AddHamsters(c echo.Context) error {
	hamster := models.Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to reading body request: %s", err))
	}

	return c.String(http.StatusOK, "we get your hamster")
}

```

```

```
