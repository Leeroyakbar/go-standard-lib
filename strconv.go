package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBool, err := strconv.ParseBool("false")
	if err != nil {
		panic(err)
	}
	fmt.Println(parseBool)

	result, err := strconv.Atoi("23")
	fmt.Println(result)

	binary := strconv.FormatInt(999, 8)
	fmt.Println(binary)
}
