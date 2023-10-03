package main

import (
	"fmt"
	"os"
)

func main() {
	/* Exercise 1.1.
	var arg string
	for _, s := range os.Args {
		arg += s + " "
	}
	fmt.Println(arg)
	*/

	for i, s := range os.Args {
		fmt.Printf("%d - %s\n", i, s)
	}
}
