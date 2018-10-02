package main

import (
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo"
	"testing"
)

func TestDummy(t *testing.T) {
 return
}

func TestExampleFunc(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET,"/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req,rec)

	xRes :=`
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
	assert.Equal(t,xRes, rec.Body.String())
}