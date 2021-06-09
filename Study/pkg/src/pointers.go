package main

import (
	"fmt"
)

func withPointer(x *int) {
	*x = *x * *x
}

type Complex struct {
	x, y int
}

func newComplex(x, y int) *Complex {
	return &Complex{x, y}
}

func main() {
	x := -2
	fmt.Printf("Address of x : %#x and Value of x :%d\n", &x, x)
	withPointer(&x)
	fmt.Printf("Address of x : %#x and Value of x :%d\n", &x, x)

	w := newComplex(4, -5)
	fmt.Println(*w)
	fmt.Println(w)
}
