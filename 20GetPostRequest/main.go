package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//fmt.Println("Creating server")
	PerformGetRequest()
	//PerformPostJsonRequest()
	//PerformPostformJsonRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// fmt.Println("Response is : ", response)
	// fmt.Println("Status is : ", response.StatusCode)
	// fmt.Println("Content length is : ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("Content is : ", content)
	fmt.Println("Content is : ", string(content))

	// another efficient method to print content get from url is

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteContent is : ", byteCount)
	fmt.Println("Content is : ", responseString.String())

}

func PerformPostJsonRequest() {
	const url = "http://localhost:8000/post"
	// json payload
	requestBody := strings.NewReader(`
	    {
			"CourseName":"GoLang",
			"Price": 0,
			"Plateform":"YouTube"
		}
	`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostformJsonRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "samridhi")
	data.Add("lastname", "sahu")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
