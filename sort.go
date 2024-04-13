package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (userSlices Users) Len() int {
	return len(userSlices)
}

func (userSlices Users) Less(i, j int) bool {
	return userSlices[i].Age < userSlices[j].Age
}

func (userSlices Users) Swap(i, j int) {
	//temp := userSlices[i]
	//userSlices[i] = userSlices[j]
	//userSlices[j] = temp

	userSlices[i], userSlices[j] = userSlices[j], userSlices[i]
}

func main() {
	users := []User{
		{"Lee", 23},
		{"Roy", 28},
		{"Akbar", 20},
	}

	sort.Sort(Users(users))

	fmt.Println(users)
}
