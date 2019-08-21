package jsonapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessWriter(t *testing.T) {
	res := httptest.NewRecorder()
	api := New(res, "")

	api.SuccessWriter(map[string]string{"status": "1"})
	assert.Equal(t, 200, res.Code)
}

func TestSuccessWriterWithCallback(t *testing.T) {
	res := httptest.NewRecorder()
	api := New(res, "ThisIsCallbackMethod")

	api.SuccessWriter(map[string]string{"status": "1"})
	assert.Equal(t, 200, res.Code)
}

func TestErrorWriter(t *testing.T) {
	res := httptest.NewRecorder()
	api := New(res, "")

	api.ErrorWriter(http.StatusForbidden, "Test Forbidden.")
	assert.Equal(t, 403, res.Code)
}
