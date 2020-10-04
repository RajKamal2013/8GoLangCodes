package main

import(
	"fmt"
	"os"
	)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file name ")
		os.Exit(1)
	}

	fileName := os.Args[1]
	destination, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create() :", err)
		os.Exit(1)
	}

	fmt.Fprintf(destination, "[%s]:", fileName)
	fmt.Fprintf(destination, "Using fmt.Fprintf in %s \n", fileName)

	defer destination.Close()
}

