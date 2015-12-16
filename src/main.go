package main

import (
	"fmt"
	"os"
)

func main() {
	println("Testing")
	if len(os.Args) != 2 { //os.Args[0] is always the path of the current running executable
		os.Exit(1)
	}

	fmt.Println("It's over", os.Args[1])
}
