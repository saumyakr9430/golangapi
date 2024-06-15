package main

import (
	"fmt"
	"sort"
)

func main() {
	//similar to vector
	var fruitlist = []string{"apple","string" , "tomaoto"}

	fmt.Println("type of fruitlist is: " , fruitlist)

	fruitlist = append(fruitlist ,"mango " , "banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highscore := make([]int , 4)
	highscore = append(highscore, 23,4,5 )
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))

	// how to remove a value of slice based on index

	var course = []string{"react" , "java" ,"go" ,"pyhon"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)







	

	
}