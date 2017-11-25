package main

import (
	"fmt"
	"net/http"
	_ "strings"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	url := "https://news.ycombinator.com/"
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}


	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b , _ := ioutil.ReadAll(resp.Body)
	status := resp.Status
	resp.Body.Close()

	fmt.Println("Status: ", status)
	fmt.Printf("%s", b)

}
