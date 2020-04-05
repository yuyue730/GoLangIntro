package main

import (
	"GoLangIntro/DistributedCrawler/config"
	"GoLangIntro/DistributedCrawler/rpcsupport"
	"GoLangIntro/DistributedCrawler/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://newcar.xcar.com.cn/m49989/",
		Parser: worker.SerializedParser{
			Name: config.ParseCarDetail,
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
