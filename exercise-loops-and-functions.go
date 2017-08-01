package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	value := 0
	for value < 100 {
		value += 1
		fmt.Println(Sqrt(float64(value)))
	}	
}