package main

import "fmt"

func main() {
	urls := []string{
    "frist",
    "second",
	}

	for index, url := range urls {
		fmt.Println(index , url)
	}
}
