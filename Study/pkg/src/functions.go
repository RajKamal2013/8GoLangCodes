package main

import (
	"fmt"
)

func unnamedMinMax(x, y int) (int, int) {
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func sort(x, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func main() {
	y := 4
	square := func(s int) int {
		return s * s
	}
	fmt.Println("The Square of ", y, "is ", square(y))

	square = func(s int) int {
		return s + s
	}

	fmt.Println("The Square of ", y, "is ", square(y))
	fmt.Println(minMax(15, 6))
	fmt.Println(namedMinMax(4, 10))
	min, max := namedMinMax(4, 10)
	fmt.Println(min, max)
}
