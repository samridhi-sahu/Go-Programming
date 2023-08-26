package main

import "fmt"

func main() {
	fmt.Println("Learning Defer here")
	defer fmt.Println("This is a defer line")
	fmt.Println("this line execute before defer line")
	fmt.Println("because this is how defer works")
}
