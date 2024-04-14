package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Lee Roy Akbar", "ee"))
	fmt.Println(strings.Split("Lee Roy; Akbar", "; "))
	fmt.Println(strings.Trim("        lee Roy       ", " "))
	fmt.Println(strings.ToLower("LEE ROy AkbAr"))
	fmt.Println(strings.ToUpper("Lee Roy Akbar"))
	fmt.Println(strings.ReplaceAll("Lee Roy", "Lee", "Lili"))
}
