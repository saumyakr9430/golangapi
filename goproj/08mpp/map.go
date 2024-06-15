package main

import (
	"fmt"

	
)

func main() {

	languages := make(map[string]string)
	
	languages["JS"] = "javascript"
	languages["go"] = "golang"
	languages["react"] = "javascript"
	languages["node"] = "javascript"



	fmt.Println(languages)

  delete(languages , "JS")
  fmt.Println(languages)

  //for loop

  for key,value := range languages{
	fmt.Printf("for key %v :   for value %v" , key , value)
  }


}