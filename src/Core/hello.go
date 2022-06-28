package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

func addNumber(a int, b int) int {
	return a + b
}

func divAndRem(a int, b int) (int, int) {
	return a / b, a % b
}

func addOne(a int) int {
	return a + 1
}
func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makeAdder(b int) func(int) int {
	return func(a int) int { return a + b }
}

func a1() {
	fmt.Println("Using Defer")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func generate_randome(a int) {
	var rdata int
	fmt.Println("Generating", a, " random numbers")
	for i := 0; i < a; i++ {
		rdata = rand.Int() % 1000
		fmt.Println("Random Data generated:", rdata)
	}
}

func main() {
	fmt.Printf("Hello Go!!!")
	// Command line arguments
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(string(os.Args[i]))
	}

	// User IO
	var num int
	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &num)
	fmt.Println("Entered number is :", num)

	reader := bufio.NewReader(os.Stdin)
	var data string
	data, _ = reader.ReadString('\n')
	fmt.Println("user added ", data)

	// File Io
	var fName string
	fName = "gotext.txt"
	fd, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Not able to create file :%s", err)
		os.Exit(1)
	}

	leng, err := fd.WriteString("Hi I am new to Go")

	if err != nil {
		log.Fatalf("Write String failed :%s", err)
		os.Exit(1)
	}
	fmt.Println("Data written in bytes :", leng)
	fmt.Println("File name ", fd.Name())
	fd.Close()

	fdata, err := ioutil.ReadFile(fName)
	fmt.Println(fdata)
	fmt.Println(string(fdata))

	fd, err = os.Open(fName)
	reader = bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Done reading file")
			break
		} else if err != nil {
			log.Fatalf("Error while reading file %s", err)
			os.Exit(1)
		}
		fmt.Println("File read : ", line)
	}
	fd.Close()

	var sname string
	sname = "Hi I am Raj Kamal and you can call me Howard Skywalker"
	for k, v := range sname {
		fmt.Println(k, ":", v, "->", string(v))
	}

	var arr [5]int

	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4

	for _, i := range arr {
		fmt.Println(i)
	}

	a := 2
	b := 3
	c := addNumber(a, b)
	fmt.Println("Adding ", a, "and ", b, "gets me:", c)

	q, r := divAndRem(2, 3)
	fmt.Println(a, "div by ", b, " results ", q, "and ", r, "as quotien and remainder ")

	y := 4
	square := func(s int) int { return s * s }
	fmt.Println("Square of ", y, " is :", square(4))

	square = func(s int) int { return s + s }
	fmt.Println("Double of ", y, " is:", square(y))

	myaddOne := addOne

	fmt.Println("adding one to ", y, "is :", myaddOne(y))

	printOperation(y, addOne)

	newAddOne := makeAdder(1)
	newAddTwo := makeAdder(2)

	fmt.Println(newAddOne(y))
	fmt.Println(newAddTwo(y))

	a1()

	generate_randome(10)
}
