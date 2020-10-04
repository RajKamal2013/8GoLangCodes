package main

import(
	"fmt"
	"time"
	)

func main()  {
	go func() {
		fmt.Println("hello ")
	}()

	go func() {
		fmt.Println("Go Routine")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}


