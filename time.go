package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2020, time.August, 2, 1, 2, 3, 4, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(valueTime)
	}
}
