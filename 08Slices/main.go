package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Slices")
	//var slice = []int{1, 3, 5, 7, 9}
	//fmt.Println("Before Elements are : ", slice)
	//fmt.Printf("Type is : %T", slice)

	// appending or inserting elements into slice at the end using append method

	//slice = append(slice, 92, 17, 4)
	//fmt.Println("After appending Elements are : ", slice)

	// slicing of slice

	//slice = append(slice[1:])
	// slice = append(slice[:3])
	// fmt.Println("After slicing Elements are : ", slice)

	scores := make([]int, 4)
	scores[0] = 23
	scores[1] = 30
	scores[2] = 13
	scores[3] = 53
	// gives run time error
	//scores[4] = 11
	//scores[5] = 33
	//scores[6] = 14
	fmt.Println("Before : ", scores)
	//scores = append(scores, 11, 33, 14)
	//fmt.Println("After : ", scores)

	// sorting method to sort slice in increasing order
	//sort.Ints(scores)
	//fmt.Println("After sorting : ", scores)

	// remove value from slices based on index
	// suppose want to remove index 2 then

	var index int = 2
	scores = append(scores[:index], scores[index+1:]...)
	fmt.Println("After removing : ", scores)
}
