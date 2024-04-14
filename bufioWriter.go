package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Lee\n")
	_, _ = writer.WriteString("Roy\n")
	writer.Flush()
}
