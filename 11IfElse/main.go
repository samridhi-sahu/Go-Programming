package main

import "fmt"

func main() {
	//fmt.Println("if else")
	age := 18
	var res string
	if age >= 21 {
		res = "Have both voting and marriage rights"
	} else if age >= 18 && age < 21 {
		res = "Have only voting right"
	} else {
		res = "Niether Have voting nor marriage rights"
	}
	fmt.Println("Result is : ", res)
}
