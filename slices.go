package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Lee", "Roy", "Akbar"}
	values := []int{100, 101, 823, 12, 23}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
}
