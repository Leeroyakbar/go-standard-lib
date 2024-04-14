package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, msg string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(msg)
	return nil
}

func addToFile(name string, msg string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(msg)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", nil
	}

	reader := bufio.NewReader(file)
	var msg string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		msg += string(line) + "\n"
	}

	return msg, nil
}

func main() {
	//createNewFile("sample.log", "this is sample log")

	addToFile("sample.log", "\nlee roy akbar")
	file, _ := readFile("sample.log")
	fmt.Println(file)
}
