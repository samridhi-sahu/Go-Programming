package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("Gorilla mux")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome</h1>"))
}
