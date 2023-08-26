package main

import "fmt"

func main() {
	//fmt.Println("Loop")
	//counting := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println("Numbers are : ", counting)

	// for loop
	// for i := 0; i < len(counting); i++ {
	// 	fmt.Println(counting[i])
	// }

	// for each loop
	// for i := range counting {
	// 	fmt.Println(counting[i])
	// }

	// for each loop with formatting
	// for index, number := range counting {
	// 	fmt.Printf("Index is %v and value at index is %v\n", index, number)
	// }

	// while loop

	// i := 1
	// for i <= 10 {
	// 	fmt.Printf("2 x %v = %v\n", i, 2*i)
	// 	i++
	// }

	// loop with break continue and goto statement

	i := 1
	for i <= 10 {
		if 2*i+1 == 7 {
			fmt.Println("Oh thats my fav number, I got it 7 hurray!")
			i++
			continue
		}
		if 2*i == 14 {
			fmt.Println("Oh thats a bad number, oops!")
			break
		}
		fmt.Printf("2 x %v = %v\n", i, 2*i)
		if i == 2 {
			goto check
		}
		i++
	}
check:
	fmt.Println("yes we did it, yay!")
}
