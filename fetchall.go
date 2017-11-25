package main

import (
	"fmt"
	"strconv"
	"io"
	"io/ioutil"
	"time"
	"net/http"
	"bufio"
	"os"
)

const (
	numPages = 3
	baseURL = "https://news.ycombinator.com/news?p="
)

func main() {
	ch := make(chan string)
	var urls []string
	for i:=1; i <= numPages; i++ {
		urls = append(urls, baseURL + strconv.Itoa(i))
	}

	var fileName string
	for idx, url := range urls {
		fileName = "./page" + strconv.Itoa(idx+1) + ".txt"
		go fetch(url, fileName, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}


func fetch(url string, fileName string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// write to file
	b, _ := ioutil.ReadAll(resp.Body)
	f, _ := os.Create(fileName)
	w := bufio.NewWriter(f)
	w.Write(b)
	f.Close()

	// check size
	nbytes, _ := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}