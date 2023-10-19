package main

import (
	"bufio"
	"os"
)

func main() {
	// lissajous(os.Stdout)
	// for _, url := range os.Args[1:] {
	// 	fetch(url)
	// }

	urls := os.Args[1:]

	fileInfo, err := os.Stdin.Stat()
	if err == nil && fileInfo.Size() > 0 {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			urls = append(urls, scan.Text())
		}
	}
	fetchall(urls)
}
