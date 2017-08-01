package main

import "fmt"

func swap(x, y string) (int, int) {
	if x == "hello" {
		return 1, 2
	} else {
		return 3, 4
	}
}

func main() {
	a, b := swap("hell", "world")
	fmt.Println(a, b)
}
