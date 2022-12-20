package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hellow World")

	// Variables
	var i int = 42 // variable name
	fmt.Printf("i = %v, type = %T\n", i, i)

	var j float32
	j = float32(1) // type convertions
	fmt.Printf("j = %v, type = %T\n", j, j)

	k := 3.1244 // direct integration without type definition
	fmt.Printf("k = %v, type = %T\n", k, k)
}
