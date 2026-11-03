package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcome to time")

	currentTime := time.Now()

	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2004 15:04:05 Monday"))

	newDate := time.Date(2024, time.June, 10, 12, 30, 0, 0, time.UTC)
	fmt.Println(newDate.Format("01-02-2024 Monday"))

}
