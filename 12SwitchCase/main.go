package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Switch Case")
	// rand.Seed(time.Now().UnixNano())
	// dice := rand.Intn(6) + 1
	// fmt.Println("Value of dice is : ", dice)

	// automatically used break by defualt this exist in golang

	// switch dice {
	// case 1:
	// 	fmt.Println("Value is 1, you can open it")
	// case 2:
	// 	fmt.Println("Value is 2")
	// case 3:
	// 	fmt.Println("Value is 3")
	// case 4:
	// 	fmt.Println("Value is 4")
	// case 5:
	// 	fmt.Println("Value is 5")
	// case 6:
	// 	fmt.Println("Value is 6, have one more chance to play")
	// default:
	// 	fmt.Println("Whats that!")
	// }

	// using fallthrough

	value := 5
	switch value {
	case 1:
		fmt.Println("Value is 1")
		fallthrough
	case 2:
		fmt.Println("Value is 2")
		fallthrough
	case 3:
		fmt.Println("Value is 3")
		fallthrough
	case 4:
		fmt.Println("Value is 4")
		fallthrough
	case 5:
		fmt.Println("Value is 5")
		fallthrough
	case 6:
		fmt.Println("Value is 6")
	default:
		fmt.Println("Whats that!")
	}

}
