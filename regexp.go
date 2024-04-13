package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`l([a-z])e`)

	fmt.Println(regex.MatchString("lee"))
}
