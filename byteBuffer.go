package main

import (
	"fmt"
	"strings"
	"sort"
)


func main() {
	// var stack bytes.Buffer
	fmt.Println(sortString("cab"))

}


func sortString(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
	sort.
}