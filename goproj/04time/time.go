package main

import (
	"fmt"
	"time"
)

func main() {

	presenTime := time.Now()

	fmt.Println(presenTime)

	fmt.Println(presenTime.Format("monday"))
}