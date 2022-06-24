package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	x := [][]string{{"morning"},{""}}
	x[1] = append(x[1], "good")
	fmt.Println(x) 
}