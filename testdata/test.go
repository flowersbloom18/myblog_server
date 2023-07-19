package main

import (
	"fmt"
	"net/url"
)

func main() {
	//encodedParam := "%E6%9D%B0%E5%85%8B%E8%9E%B3%E8%9E%82"
	encodedParam := "test1"
	decodedParam, err := url.QueryUnescape(encodedParam)
	if err != nil {
		fmt.Println("URL decoding error:", err)
		return
	}
	fmt.Println("Decoded parameter:", decodedParam)
}
