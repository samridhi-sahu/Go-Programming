package main

import "fmt"

func main() {
	fmt.Println("Array")
	// var myarray [5]int
	// myarray[0] = 1
	// myarray[1] = 3
	// myarray[2] = 5
	// myarray[3] = 7
	// myarray[4] = 9

	var myarray = [5]int{1, 3, 5, 7, 9}
	fmt.Println("Elements are : ", myarray)
	fmt.Println("Length is : ", len(myarray))
}
