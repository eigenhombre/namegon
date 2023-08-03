package main

import (
	"fmt"
)

func main() {
	CHAINLEN := 2
	ch := buildChain(MaleIshNames, CHAINLEN)
	for i := 0; i < 10; i++ {
		fmt.Print(generateName(ch, CHAINLEN))
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println("OK")
}
