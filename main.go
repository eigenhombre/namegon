package main

import (
	"fmt"
)

func main() {
	mkman := Namer(MaleIshNames, 4)
	for i := 0; i < 10; i++ {
		fmt.Print(mkman())
		fmt.Print(" ")
	}
	fmt.Println()
	fmt.Println("Done.")
}
