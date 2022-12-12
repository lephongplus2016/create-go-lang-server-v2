package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	catJson = `{
		"name": "kity",
		"type": "afica cat"
	}`
	catString   = `we get your cat`
	yalloString = `yallo from the webside`
)

func TestAddCat(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/cats", strings.NewReader(catJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// log
	// fmt.Println("rec:", rec)
	// fmt.Println("rec.Code:", rec.Code)
	// fmt.Println("rec.Body.String():", rec.Body.String())

	// Assertions
	if assert.NoError(t, AddCat(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, catString, rec.Body.String())
	}
}

func TestYallo(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	// Assertions
	if assert.NoError(t, Yallo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, yalloString, rec.Body.String())
	}
}

func TestGetCats(t *testing.T) {
	// Setup
	e := echo.New()
	// query params
	q := make(url.Values)
	q.Set("name", "alison")
	q.Set("type", "asia")
	req := httptest.NewRequest(http.MethodGet, "/cats?"+q.Encode(), nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// use param
	c.SetPath("/cats/:data")
	c.SetParamNames("data")
	c.SetParamValues("json")

	// Assertions
	if assert.NoError(t, GetCats(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t,
			`{"cat":"alison","type":"asia"}
`, rec.Body.String())
	}
}

// in case wrong
func TestGetCatsMissingDataInQueryParam(t *testing.T) {
	// Setup
	e := echo.New()
	// query params
	q := make(url.Values)
	q.Set("name", "alison")
	q.Set("type", "asia")
	req := httptest.NewRequest(http.MethodGet, "/cats?"+q.Encode(), nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// use param
	c.SetPath("/cats/:data")
	c.SetParamNames("data")
	c.SetParamValues("a")

	// Assertions
	if assert.NoError(t, GetCats(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t,
			`{"error":"you need to let's us know if you want to json or string data"}
`, rec.Body.String())
	}
}
