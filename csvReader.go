package main

import (
	"container/list"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func readerIO(csvStr string) list.List {
	reader := csv.NewReader(strings.NewReader(csvStr))

	var records list.List

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		records.PushBack(record)
		next := records.Front()
		next = next.Next()

	}
	return records
}

func main() {
	csvString := "lee,roy,akbar\n" +
		"lili,rahma,yani"

	var records list.List = readerIO(csvString)
	fmt.Println(records.Front().Value)
	fmt.Println(records.Front().Next().Value)

}
