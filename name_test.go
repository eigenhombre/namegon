package namegon_test

import (
	"testing"

	"github.com/eigenhombre/namegon"
)

func TestNameSmokeTest(t *testing.T) {
	result := namegon.Namer(namegon.ExampleMaleNames, 5)()
	if len(result) < 5 {
		t.Errorf("Expected name to be at least 5 characters, got %s", result)
	}
}
