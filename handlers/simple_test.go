package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChrisTheShark/golang-interview/handlers"
	"github.com/stretchr/testify/assert"
)

func TestSimpleHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/simple", nil)
	w := httptest.NewRecorder()

	handlers.SimpleHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "Hello, World!", w.Body.String())
}
