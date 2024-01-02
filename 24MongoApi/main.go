package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samridhi-sahu/mongoapi/router"
)

func main() {
	fmt.Println("learn mongo api")
	r := router.Router()
	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000...")
}
