package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome in Go Programming "
	reader := bufio.NewReader(os.Stdin)
	// comma ok syntax
	//input, _ := reader.ReadString('\n')

	// comma error syntax
	input, err := reader.ReadString('\n')
	fmt.Print(welcome, input)
}
