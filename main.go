package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildChain([]string{"f",
		"fo", "fooo", "foo", "bar", "baz", "soor", "zoom", "zoot"}, 2))
	// fmt.Println(buildChain(MaleIshNames, 2))
	fmt.Println("OK")
}
