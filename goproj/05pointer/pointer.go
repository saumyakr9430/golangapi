package main

import "fmt"

//refrence to memory address
// pointer gives us the suriety that a specific value is returned not its copy(it gives sometimes error)
func main() {

	// var ptr *int

	// fmt.Println("vlaue of pointer is:" , &ptr )

	number := 23

	var ptr = &number

	*ptr = *ptr * 2

	fmt.Println("value:", *ptr) // value as * is before

	fmt.Println("address:", &number) // address

}
