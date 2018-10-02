package main

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	return
}

func TestHealthCheckFun(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, healthCheck(c))
	assert.Equal(t, "ok!", rec.Body.String())
}

func TestExampleFunc(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	xRes := `
		<code>
			Protocol: HTTP/1.1<br>
			Host: example.com<br>
			Remote Address: 192.0.2.1:1234<br>
			Method: GET<br>
			Path: /<br>
			Message : Test
		</code>
	`

	assert.NoError(t, example(c))
	assert.Equal(t, xRes, rec.Body.String())
}
