package main

import (
"os"
"fmt"
"strconv"
)

func main() {
	arguments := os.Args
	sum := 0
	for i := 0; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i])
		sum = sum + temp
	}
	fmt.Println("Sum : ", sum)
}