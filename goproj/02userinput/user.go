package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("what is to be input")

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks", input)

}
