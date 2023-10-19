package main

func main() {
	// lissajous(os.Stdout)
	// for _, url := range os.Args[1:] {
	// 	fetch(url)
	// }

	urls := getUrls()
	fetchall(urls)
}
