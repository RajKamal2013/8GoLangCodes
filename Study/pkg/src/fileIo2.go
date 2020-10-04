package main

import(
	"fmt"
	"bufio"
	"os"
	)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please Provide A fileName")
		os.Exit(1)
	}

	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error Opening %s :%s", fileName, err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Err() != nil {
			fmt.Printf("Error Reading file %s ", err)
			os.Exit(1)
		}
		fmt.Println(line)
	}

	defer f.Close()
}



