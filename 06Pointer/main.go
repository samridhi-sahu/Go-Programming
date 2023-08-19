package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is : ", ptr)
	myNumber := 11
	var ptr = &myNumber
	fmt.Println("Address which pointer is referring to : ", ptr)
	fmt.Println("Value at address which pointer is referring to : ", *ptr)
	fmt.Println("Before : ", myNumber)
	*ptr = *ptr + 1
	fmt.Println("After : ", myNumber)
}
