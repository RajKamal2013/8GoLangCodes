package main

import(
	"fmt"
	"os"
	"io"
	)

func countChars(r io.Reader) int {
	buf :=  make([]byte, 16)
	total := 0

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}

		if err == io.EOF {
			break;
		}

		total = total + n
	}

	return total
}

func writeNumberOfChars(w io.Writer, x int) {
	fmt.Fprintf(w, "%d\n", x)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide the filename")
		os.Exit(1)
	}

	fileName := os.Args[1]
	_, err := os.Stat(fileName)

	if err != nil {
		fmt.Println("Cannot open file :", err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open the file :", err)
		os.Exit(1)
	}
	defer f.Close()

	chars := countChars(f)

	fileName = fileName + ".count"
	f, err = os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create() :", err)
		os.Exit(1)
	}

	defer f.Close()
	writeNumberOfChars(f, chars)
}
