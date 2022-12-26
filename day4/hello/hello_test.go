package hello

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/hello", nil)

	helloHandler(rec, request)
	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "hello")
}
