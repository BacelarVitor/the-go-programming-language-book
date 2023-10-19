package main

import (
	"fmt"
	"time"
)

func fetchall(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
