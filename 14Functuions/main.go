package main

import "fmt"

func main() {
	fmt.Println("Functions")
	// void function without arguments
	greet()

	// void function with arguments
	greetWithCountry("India")

	// return type as int function with arguments
	result := addition(5, 7)
	fmt.Println("Sum of two number is : ", result)

	// function with multiple arguments when we dont know how many arguments need to pass
	Proresult := ProAddition(2, 4, 6, 8, 5, 7)
	fmt.Println("Sum of numbers are : ", Proresult)

	// function with two return type
	ProAns, msg := Prosum(2, 4, 5, 7)
	fmt.Println("Sum of numbers are : ", ProAns)
	fmt.Println("My Message is : ", msg)

}

func greet() {
	fmt.Println("Jai Hind")
}

func greetWithCountry(country string) {
	fmt.Println("Jai Hind", country)
}

func addition(num1 int, num2 int) int {
	return num1 + num2
}

func ProAddition(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func Prosum(numbers ...int) (int, string) {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum, "Hurry! we did it."
}
