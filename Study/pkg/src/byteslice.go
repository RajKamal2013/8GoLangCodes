package main

import (
	"fmt"
	"os"
	"io/ioutil"
	)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	fileName := os.Args[1]

	aByteSlice := []byte("Mr Raj Kamal is learning GoLang")
	ioutil.WriteFile(fileName, aByteSlice, 0644)

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	anotherByteSlice := make([]byte, 100)
	n, err := f.Read(anotherByteSlice)
	fmt.Printf("Read %d bytes : %s ", n, anotherByteSlice)
}
