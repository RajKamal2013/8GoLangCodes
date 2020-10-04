package main

import(
	"fmt"
	"sort"
)

type astruct struct {
	person  string
	height  int
	weight  int
}

func main() {
	mySlice := make([]astruct, 0)
	a := astruct{"Raj Kamal",  167,  65}
	mySlice =  append(mySlice, a);
	a = astruct{"Howard Skywalker",  167,  63}
	mySlice =  append(mySlice, a);
	a = astruct{"Sonu ",  167,  66}
	mySlice =  append(mySlice, a);

	fmt.Println("0:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
					return mySlice[i].weight < mySlice[j].weight
				})
	fmt.Println(">:",  mySlice);
	sort.Slice(mySlice, func(i, j int) bool {
					return mySlice[i].weight > mySlice[j].weight
				})
	fmt.Println("<:",  mySlice);
}
