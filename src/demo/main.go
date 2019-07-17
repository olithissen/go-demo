package main

import (
	adder "demo/adder"
	"fmt"
)

func main() {
	fmt.Println("Hello gitpod.io!")
	result := adder.Add(1, 2)
	fmt.Printf("%d", result)
}
