package main

import (
	"fmt"
	"math"
)

func squareCube(i int) {
	for num := 1; num <= i; num++ {
		x := math.Cbrt(float64(num))
		y := math.Sqrt(float64(num))

		cube := x == float64(int(x))
		square := y == float64(int(y))

		switch {
		case cube == true && square == true:
			fmt.Println("SquareCube")
		case cube == true && square == false:
			fmt.Println("Cube")
		case cube == false && square == true:
			fmt.Println("Square")
		default:
			fmt.Println(num)
		}
	}
}

func main() {
	squareCube(100)
}
