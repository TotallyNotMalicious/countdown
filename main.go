package main

import (
	"fmt"
	"time"
)

func main() {
	var input int

	fmt.Print("Time To Count Down: ")
	fmt.Scan(&input)
	fmt.Println()

	for i := input; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("Done.")
}
