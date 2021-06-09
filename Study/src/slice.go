package main

import (
	"fmt"
)

func change(x []int) {
	x[3] = -2
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}

func main() {
	aSlice := []int{-1, 4, 5, 0, 7, 9}
	fmt.Printf("Before Change :")
	printSlice(aSlice)
	change(aSlice)
	fmt.Printf("After Change :")
	printSlice(aSlice)
	fmt.Printf("Before cap = %d, Length = %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, 100)
	fmt.Printf("After cap = %d, Length = %d\n", cap(aSlice), len(aSlice))
	printSlice(aSlice)
	anotherSlice := make([]int, 4)
	fmt.Printf("New Slice with all zeros ")
	printSlice(anotherSlice)
}
