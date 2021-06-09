package main

import "fmt"

//import "os"

func unNamedMinMax(x, y int) (int, int) {
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

func main() {
	/* Printing */
	fmt.Printf("hello, world\n")
	/* Input and output with Variable declaration */
	var num int
	fmt.Printf("Enter the number")
	fmt.Scanln(&num)
	fmt.Printf("Number entered is :%d \n", num)
	/* Function */
	min, max := unNamedMinMax(15, 19)
	fmt.Printf("Max and min :%d, %d \n", max, min)
	/* Anonymous function */
	square := func(s int) int {
		return s * s
	}
	var y int
	y = 4
	fmt.Printf("The square of %d is %d \n", y, square(y))
	square = func(s int) int {
		return s + s
	}
	fmt.Printf("The sum of %d is %d\n", y, square(y))
	/* array */
	arr := [4]int{12, 45, 90, -1}
	for _, number := range arr {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")
	/* slice */
	aslice := []int{23, 89, 90}
	for _, number := range aslice {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")

	aslice = append(aslice, 91)
	for _, number := range aslice {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")
	fmt.Printf("Size and Capacity : %d and %d \n", len(aslice), cap(aslice))
	anotherSlice := make([]int, 4)
	anotherSlice = append(anotherSlice, 9)
	anotherSlice = append(anotherSlice, 9)
	anotherSlice = append(anotherSlice, 9)
	for _, number := range anotherSlice {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")
	/* Loop */
	i := 1
	var sum int
	for sum = 1; sum < 10; sum++ {
		i = i + sum
		fmt.Println(i)
	}
	fmt.Println(i)
	/* while loop => for */
	sum = 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

}
