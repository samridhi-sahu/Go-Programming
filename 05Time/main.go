package main

import (
	"fmt"
	"time"
)

func main() {
	// presentTime := time.Now()
	// fmt.Print("Full Time : ")
	// fmt.Println(presentTime)

	// want to format and print date only
	// so for that we have to do formatting with use of standard date format

	// fmt.Print("Time : ")
	// fmt.Println(presentTime.Format("15:04:05"))
	// fmt.Print("Date : ")
	// fmt.Println(presentTime.Format("01-02-2006"))
	// fmt.Print("Day : ")
	// fmt.Println(presentTime.Format("Monday"))
	// fmt.Print("Date Day Time : ")
	// fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	// create time manually
	createDate := time.Date(2001, time.September, 11, 05, 30, 00, 00, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("Monday"))
}
