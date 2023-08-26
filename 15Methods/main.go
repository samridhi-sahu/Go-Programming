package main

import "fmt"

func main() {
	//fmt.Println("Methods")
	sam := Users{"Samridhi", "abs@gmail.com", true, 23}

	// print user
	//fmt.Println("User is : ", sam)

	// accessing only name from user struct
	fmt.Printf("Email is %+v\n", sam.Email)
	sam.GetStatus()
	sam.NewEmail()
	fmt.Printf("%+v\n", sam.Email)
}

type Users struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u Users) GetStatus() {
	fmt.Println("is user active : ", u.Status)
}

func (u Users) NewEmail() {
	u.Email = "abcd@gmail.com"
	fmt.Println("New email is : ", u.Email)
}
