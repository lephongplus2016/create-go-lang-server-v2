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
