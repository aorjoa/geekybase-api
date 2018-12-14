package main

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestResponseStructureOfGetFeeds(t *testing.T) {
	expectedResponse := `{
	   "version": "https://jsonfeed.org/version/1",
	   "title": "jmoiron.net blog",
	   "home_page_url": "http://jmoiron.net/blog",
	   "description": "discussion about tech, footie, photos",
	   "author": {
	     "name": "Jason Moiron"
	   },
	   "items": [
	     {
	       "id": "",
	       "url": "http://jmoiron.net/blog/limiting-concurrency-in-go/",
	       "title": "Limiting Concurrency in Go",
	       "summary": "A discussion on controlled parallelism in golang",
	       "date_published": "2018-12-14T09:09:30.88107+07:00",
	       "author": {
	         "name": "Jason Moiron"
	       }
	     },
	     {
	       "id": "",
	       "url": "http://jmoiron.net/blog/logicless-template-redux/",
	       "title": "Logic-less Template Redux",
	       "summary": "More thoughts on logicless templates",
	       "date_published": "2018-12-14T09:09:30.88107+07:00"
	     },
	     {
	       "id": "",
	       "url": "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/",
	       "title": "Idiomatic Code Reuse in Go",
	       "summary": "How to use interfaces <em>effectively</em>",
	       "date_published": "2018-12-14T09:09:30.88107+07:00"
	     }
	   ]
	 }`

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/feeds", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := getFeeds(c)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedResponse, rec.Body.String())

}
