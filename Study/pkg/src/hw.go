package main

import "fmt"
import "os"

func main() {
	argument := os.Args
	for i := 0; i < len(argument); i++ {
		fmt.Println(argument[i])
	}
}
