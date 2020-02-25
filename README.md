# Introduction into Go Programming Language

## Installation and Configuration

  * Official website: https://golang.org/. And download the installation package from the website and install.
  * Confirm `go` has been successfully installed by typing in
    ```
    Yus-MacBook-Pro:~ yyu196$ go version
    go version go1.13.7 darwin/amd64
    ```
  * Turn on `GO111MODULE` and install `golangimports` dependency by typing in
    ```
    Yus-MacBook-Pro:~ yyu196$ go env -w GO111MODULE=on
    Yus-MacBook-Pro:~ yyu196$ go get -v golang.org/x/tools/cmd/goimports
    ```
  * Install all `go` related extensions in Visual Studio Code.
  * Create `go.mod` file under `./` by typing in
    ```
    Yus-MacBook-Pro:GoLangIntro yyu196$ go mod init FundamentalGrammer
    ```
    Suppose we have a simple helloworld go file called `FundamentalGrammer` directory, and we can run the file under `./` which contains `go.mod` by typing
    ```
    Yus-MacBook-Pro:GoLangIntro yyu196$ go run FundamentalGrammer/basic.go 
    Hello World
    ```
  * Setup proxy in Mainland China. Go to https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md. Default proxy is `GOPROXY="https://proxy.golang.org,direct"`. Type in 
    ```
    go env -w GOPROXY=https://goproxy.cn,direct
    ```
  * Install gin framework by entering command `go get -u github.com/gin-gonic/gin`.
  * Install zap library entering command `go get -u go.uber.org/zap`


## Project: Implement a `go` web crawler on an car information website
### Milestone 1: Single Thread web crawler
  * Create a `/SingleThreadCrawler` directory to store source codes.
  * Fetch and decode HTML from `http://newcar.xcar.com.cn/`. 
    * Install Go Text library by entering `go get -u golang.org/x/text` in the command line.
    * The Kanji is in a wrong encoding way, we need to do conversion. Call `transform.NewReader` on the original response body to Convert its decoder from `GBK` to `UTF-8`.
    * For code scalable reason, we also need to install Go Net library by entering `go get -u golang.org/x/net`. This library offers a functionality to detect the decoder from an html text.
    * Create a new `determineEncoding` that takes in an response body `io.Reading` and return `encode.Encoding` that includes the decoder format.
  * Go over html text and extract all Car model ids to append to `http://newcar.xcar.com.cn/` as the next url
  * Diagrams of the system
    ```
                       **********
                       * Parser *
                       **********
                            |
             Text -> Parser | Requests, Items -> Engine
                            |
    ******** request-> **********  URL-> ***********
    * Seed *-----------* Engine *--------* Fetcher *
    ********           ********** <-Text ***********
                            |
                            |
                     **************
                     * Task Queue *
                     **************
    ```
  * `main()` in Main.go passes `http://newcar.xcar.com.cn/` into the engine's `Run()` to start the crawler application. Logic of crawling is shown in the diagram below
    ```
             -------------| 
             |            | OtherListUrls
    *****************     |         ******************                  ***************
    * Car List Page *---------------* Car Model Page *------------------* Car Details *
    *****************  ModelUrls->  ****************** CarDetailsUrls-> ***************
    ```