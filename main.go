package main

import (
	"fmt"
	"os"
)

func main() {
	var arg string
	for _, s := range os.Args {
		arg += s + " "
	}
	fmt.Println(arg)
}
