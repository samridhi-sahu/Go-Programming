package main

import "fmt"

func main() {
	fmt.Println("Maps")
	subjects := make(map[string]string)
	subjects["os"] = "operating system"
	subjects["cn"] = "computer network"
	subjects["ds"] = "data structure"
	subjects["cd"] = "compiler design"
	fmt.Println("Values inside the map : ", subjects)
	// note : this print map with key value pair in increasing order,
	//coz when we intialize keys it will store in memory in increasing order

	// accessing value using key only
	//fmt.Println("Value at key ds is : ", subjects["ds"])

	// deleting key value pair from map using delete method

	delete(subjects, "cd")
	fmt.Println("Now after deleting Values inside the map : ", subjects)
}
