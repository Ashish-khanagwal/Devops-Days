package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time study of GoLang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2019, time.March, 19, 12, 0, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
