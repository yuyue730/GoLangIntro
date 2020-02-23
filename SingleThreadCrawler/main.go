package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://newcar.xcar.com.cn/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code = %d", resp.StatusCode)
		return
	}

	// Convert `resp.Body` from GBK format to UTF-8 Format so that the Kanji
	// character can be idenfied by calling `ioutil.ReadAll` next
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
