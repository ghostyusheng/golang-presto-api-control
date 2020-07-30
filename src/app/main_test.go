package main

import (
	"fmt"
	"mymath"
)

type man struct {
	name string
	age  int
}

func kath() {
	fmt.Println("test")
}

func main() {
	//math()
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
