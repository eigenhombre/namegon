package main

import (
	"testing"
)

func TestName(t *testing.T) {
	generateName(buildChain([]string{"aa", "bb"}, 2), 2)
}
