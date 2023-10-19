package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const filePath = "featchingData.txt"

func fetchall(urls []string) {
	flag.CommandLine.Set("align", "10")
	start := time.Now()
	ch := make(chan string)

	var fileContent string
	fileContent += fmt.Sprintf("\n###### %s ######\n\n", start.Format(time.RFC822))
	fileContent += fmt.Sprintln("time\t bytes\t url\n")

	for _, url := range urls {
		go fetchGo(url, ch)
	}
	for range urls {
		fileContent += fmt.Sprintln(<-ch)
	}
	fileContent += fmt.Sprintf("\n%.2fs\t elapsed\n", time.Since(start).Seconds())

	writeFile(fileContent)
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
	ch <- fmt.Sprintf("%.2fs %7d    %s", secs, nbytes, url)

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

func writeFile(writableContent string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.WriteString(file, writableContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = file.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte(writableContent))
	if err != nil {
		// Handle the error.
		fmt.Println(err)
		return
	}

	_, err = file.Write(buffer[:n])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
