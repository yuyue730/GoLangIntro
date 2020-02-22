package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	fmt.Println("Go Programming Language Usage of http client")
	// Naive way of requesting imooc.com
	// resp, err := http.Get("http://www.immoc.com/")

	// Setup request change header configuration and request bing.com
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com/", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")

	// resp, err := http.DefaultClient.Do(request)
	// Check redirect using no default client
	fmt.Println("1. Print out redirect.")
	client := http.Client{
		CheckRedirect: func(
			req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	fmt.Println()
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("2. Dump response to console")
	fmt.Printf("Body = %s\n", s)
}
