package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"lee", "roy", "akbar"})
	_ = writer.Write([]string{"lili", "roy", "akbar"})

	writer.Flush()
}
