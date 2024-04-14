package main

import (
	"encoding/base64"
	"fmt"
)

func encode(str string) string {
	var encoded = base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

func decode(str string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return decoded
}

func main() {
	encoded := encode("Lee Roy Akbar")
	fmt.Println(encoded)

	decoded := decode(encoded)
	fmt.Println(string(decoded))
}
