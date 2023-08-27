package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "Hey this is my first file content"

	// creating a file with name firstfile
	file, err := os.Create("./firstfile.txt")
	if err != nil {
		panic(err)
	}

	// writing something onto the file
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is : ", length)
	defer file.Close()

	// reading something from the file
	data, err := ioutil.ReadFile("./firstfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Data inside the file is that : ", string(data))
}
