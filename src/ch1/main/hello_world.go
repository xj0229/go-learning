package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Now Args length 0 is: " + os.Args[0])
		fmt.Println("Now Args length 1 is: " + os.Args[1])
	}
	fmt.Println("Hello world")
	os.Exit(-1)
}
