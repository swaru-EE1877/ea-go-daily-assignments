package hello_world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnHelloWorld(t *testing.T) {
	greeting := HelloWorld()

	assert.Equal(t, greeting, "Hello, world.")
}
