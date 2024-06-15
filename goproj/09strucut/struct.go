package main

import "fmt"

func main() {
	// classes in golang

	hitesh := User{"saumya", "email.go", true, 16}
	fmt.Println(hitesh)
	//for structure
	fmt.Printf("hitesh details are: %+v\n" , hitesh)

}

type User struct {
	//first letter capital for access
	Name   string
	Email  string
	Status bool
	Age    int
}
