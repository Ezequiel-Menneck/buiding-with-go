package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan string)
	go fetchallex("https://api.stackexchange.com/2.3/questions?order=desc&sort=activity&site=stackoverflow&pagesize=100", ch)
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			return
		}
	}
}

func fetchallex(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create("response.txt")
	if err != nil {
		ch <- fmt.Sprintf("while creating archive %s: %v", url, err)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
