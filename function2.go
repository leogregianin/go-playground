package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func main() {
	fmt.Println(add(42.1, 13.2))
}
