package main

import "fmt"

func main() {
	// fmt.Println("Learning Defer here")
	// defer fmt.Println("This is a defer line")
	// fmt.Println("this line execute before defer line")
	// fmt.Println("because this is how defer works")

	// if multiple defer statment thier then they will deferred in reverse order in LIFO fashion

	fmt.Println("Hello")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("world")
	mydefer()
}

//logic : hello->world->mydefer()=> 0 1 2 3 4 (print in reverse)->(one->two->three (print in reverse))
//output : hello world 4 3 2 1 0 three two one

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
