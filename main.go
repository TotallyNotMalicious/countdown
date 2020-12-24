package main

import (
	"fmt"
	"time"
)

func main() {
	var input int

	fmt.Print("Number Of Seconds To Count Down: ")
	fmt.Scan(&input)
	fmt.Println()

	for i := input; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("Done.")
}
