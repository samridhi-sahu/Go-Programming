package main

import "fmt"

func main() {
	fmt.Println("Structs")
	sam := Users{"Samridhi", "abs@gmail.com", true, 23}

	// print user
	fmt.Println("User is : ", sam)

	// print user with formatting
	fmt.Printf("%+v\n", sam)

	// accessing only name from user struct
	fmt.Printf("Name is %v", sam.Name)
}

type Users struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
