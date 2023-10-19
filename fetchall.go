package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetchall(urls []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		fmt.Printf("featching %s\n", url)
		go fetchGo(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchGo(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

func getUrls() []string {
	urls := os.Args[1:]
	fileInfo, err := os.Stdin.Stat()
	if err == nil && fileInfo.Size() > 0 {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			urls = append(urls, scan.Text())
		}
	}
	return urls
}
