package main

import (
	rpcdemo "GoLangIntro/FundamentalGrammer/Rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	fmt.Println("Go language Rpc Client Demo")
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var result float64
	client := jsonrpc.NewClient(conn)
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	fmt.Printf("10/3 = %v, err = %v\n", result, err)
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	fmt.Printf("10/0 = %v, err = %v", result, err)
}
