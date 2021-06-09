package main

import (
	"fmt"
	"strconv"
)

func main() {
	anArray := [4]int{1, 2, -14, 0}
	aMap := make(map[string]int)
	aMap["Mon"] = 0
	aMap["Tue"] = 1
	aMap["Wed"] = 2
	aMap["Thu"] = 3
	aMap["Fri"] = 4
	aMap["Sat"] = 5
	aMap["Sun"] = 6

	fmt.Printf("Sunday is the %d day of week \n", aMap["Sun"])

	_, ok := aMap["Tuesday"]

	if ok {
		fmt.Printf("No Tuesday ")
	} else {
		fmt.Printf("Yes Tuesday")
	}

	count := 0
	for key, _ := range aMap {
		count++
		fmt.Printf("\n")
		fmt.Printf("The Map has %s Key \n", key)
	}
	fmt.Printf("The Map has %d elements \n", count)

	count = 0
	delete(aMap, "Fri")

	for key, _ := range aMap {
		count++
		fmt.Printf("\n")
		fmt.Printf("The Map has %s Key \n", key)
	}
	fmt.Printf("The Map has %d elements \n", count)

	anotherMap := map[string]int{
		"one":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
	}
	anotherMap["Six"] = 6
	count = 0
	for key, _ := range anotherMap {
		count++
		fmt.Printf("\n")
		fmt.Printf("The Map has %s Key \n", key)
	}
	fmt.Printf("The AnotherMap has %d elements \n", count)

	newMap := make(map[string]int)
	length := len(anArray)

	for i := 0; i < length; i++ {
		fmt.Printf("%s ", strconv.Itoa(i))
		newMap[strconv.Itoa(i)] = anArray[i]
	}

	for key, value := range newMap {
		fmt.Printf("Key : %s Value : %d \n", key, value)
	}
}
