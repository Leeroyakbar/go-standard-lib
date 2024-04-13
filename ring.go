package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 1; i <= data.Len(); i++ {
		data.Value = "Nilai ke " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(val any) {
		fmt.Println(val)
	})
	fmt.Println(data.Value)
}
