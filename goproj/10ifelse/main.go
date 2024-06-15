package main

import "fmt"

func main() {
	login := 23

	if login < 10 {
		fmt.Println("less than 10")
	}else if login <20{
		fmt.Println("less than 20")
	}else{
		fmt.Println("greater than 10 & 20")
	}

}