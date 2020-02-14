package main

import (
	"GoLangIntro/FundamentalGrammer/Retriever/mock"
	"GoLangIntro/FundamentalGrammer/Retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "Yue",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake Hello World",
	})
	return s.Get(url)
}

func main() {
	fmt.Println("Go language Interface with Retriever as an example")
	fmt.Println("1. Print mock retriever content.")
	mockRetriever := mock.Retriever{Contents: "This is a mock retriever"}
	rMock := &mockRetriever
	inspectRetriever(rMock)
	fmt.Println()

	fmt.Println("2. Print real retriever content.")
	var rReal Retriever
	rReal = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspectRetriever(rReal)
	fmt.Println()

	fmt.Println("3. Type assertion on real retriever")
	if mockRetriever, ok := rReal.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever")
	}
	fmt.Println()

	fmt.Println("4. Try a session")
	fmt.Println("Session return = ", session(rMock))
}

func inspectRetriever(r Retriever) {
	fmt.Printf("r Type = %T; r = %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents = ", v.Contents)
	case *real.Retriever:
		fmt.Println("User Agent = ", v.UserAgent)
	}
}
