package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web request handling")
	const url = "https://samridhi-sahu.github.io/Portfolio/"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is a type of %T\n", response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// return html of website
	fmt.Println("Data is : ", string(data))
}
