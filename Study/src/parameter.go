package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	minusI := false

	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			minusI = true
			break
		}
	}

	if minusI {
		fmt.Println("Got the -i Parameter ")
		fmt.Println("y/n: ")
		var answer string
		fmt.Scanln(&answer)
		fmt.Println("You Entereed : ", answer)
	} else {
		fmt.Println("The -i is not set ")
	}
}
