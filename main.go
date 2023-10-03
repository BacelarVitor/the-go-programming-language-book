package main

func main() {
	/* Exercise 1.1.
	var arg string
	for _, s := range os.Args {
		arg += s + " "
	}
	fmt.Println(arg)
	* Exercise 1.2
	for i, s := range os.Args {
		fmt.Printf("%d - %s\n", i, s)
	}
	* Exercise 1.3
	startJoin := time.Now()
	fmt.Print(strings.Join(os.Args[1:], " "))
	endJoin := time.Now()
	executionJoin := endJoin.Sub(startJoin)

	startFor := time.Now()
	var arg string
	for _, s := range os.Args {
		arg += s + " "
	}
	fmt.Println(arg)
	endFor := time.Now()
	executionFor := endFor.Sub(startFor)

	fmt.Println(executionJoin)
	fmt.Println(executionFor)
	*/
}
