package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("user input:")

	reader := bufio.NewReader(os.Stdin)

	input , _:= reader.ReadString('\n')

	fmt.Println("the user input is:" , input)

	newrating , err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
        fmt.Println("error has occured: " , err)
	} else{

		fmt.Println("new rating is:", newrating+1)
	}

	


}
