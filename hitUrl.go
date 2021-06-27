package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	//empty map 초기화. 아니면 err
	var results = make(map[string]string)
	urls := []string{
		"https://www.youtube.com/",
		"https://twitter.com/",
		"https://www.twitch.tv/",
		"https://www.pixiv.net/",
		"https://www.amazon.com/",
		"https://google.com/",
		"https://reddit.com/",
		"https://github.com/",
		"https://facebook.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}
func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	//400 부터 문제가있어서..
	if err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}
	return nil
}
