package main

import (
	"testing"
)

func TestNameSmokeTest(t *testing.T) {
	generateName(buildChain([]string{"aa", "bb"}, 2), 2)
}
