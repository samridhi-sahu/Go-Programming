package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Gives Rating")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("your ratings type is : %T", input)
	// conversion of string into float 64
	newInput, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Printf("Now your ratings type is : %T", newInput)
}
