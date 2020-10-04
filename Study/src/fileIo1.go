package main

import (
	"fmt"
	"io"
	"os"
	)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error Opening %s:%s", fileName, err)
		os.Exit(1)
	}

	buf := make([]byte, 16)
	if _, err := io.ReadFull(f, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}

	io.WriteString(os.Stdout, string(buf))
	fmt.Println()
	defer f.Close()
}
