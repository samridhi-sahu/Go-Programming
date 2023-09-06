package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.google.com/search?q=html+css+javascript&rlz=1C1CHBD_enIN970IN970&oq=&aqs=chrome.3.35i39i362j46i39i362j35i39i362l4j46i39i199i362i465j35i39i362.3455j0j15&sourceid=chrome&ie=UTF-8"

func main() {
	//fmt.Println("Handling URL")
	//fmt.Println(myUrl)

	//result, _ := url.Parse(myUrl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	//qparams := result.Query()

	// fmt.Printf("%T\n", qparams)
	// fmt.Println(qparams)
	// fmt.Println(qparams["sourceid"])

	// for _, val := range qparams {
	// 	fmt.Println("Param is : ", val)
	// }

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "samridhi-sahu.github.io",
		Path:   "/Portfolio/",
	}
	newUrl := partsOfUrl.String()
	fmt.Println("New URL is : ", newUrl)
}
