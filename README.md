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
  
  * Diagrams of the system

    <img src="./Images/OverviewDiagram.png" height=70% width=70%>

  * Fetcher implemetation details
    * Install Go Text library by entering `go get -u golang.org/x/text` in the command line.
    * The Kanji is in a wrong encoding way, we need to do conversion. Call `transform.NewReader` on the original response body to Convert its decoder from `GBK` to `UTF-8`.
    * For code scalable reason, we also need to install Go Net library by entering `go get -u golang.org/x/net`. This library offers a functionality to detect the decoder from an html text.
    * Create a new `determineEncoding` that takes in an response body `io.Reading` and return `encode.Encoding` that includes the decoder format.
  * Parser Implementation Diagram
  
    <img src="./Images/ParserLogic.png" height=70% width=70%>

  * Wrap `parser` functionality as a struct in `engine/types`. Create the struct with parse function object and the name of the function. Expose the method to parse contents with Items and more urls in the parse function return.

### Milestone 2: Concurrent web crawler
  * Merge functionality of Fetcher and Parser into a worker function in `engine`. 
  * For concurrent web crawler, we will simplement a scheduler that schedule `worker`s in the `engine`'s `Run` function.
  * `engine` will send `Request`s to scheduler and scheduler will coordinate `worker`s to send request and parse information. Please refer to the diagram below.

    <img src="./Images/ConcurrentCrawlerArchitecture.png" height=40% width=40%>

  * A simple Scheduler will create a `goroutine` for each Request and have a single worker to act on all `goroutine`s. Please refer to the diagram below.

    <img src="./Images/SimpleScheduler.png" height=40% width=40%>

  * A Queued Scheduler set up two queues one for worker and the other for request. When a new worker or request comes it, it will add that worker or request item at the back of the queue. When we need the worker to work on the request, we pop both front item from Request and Worker Queue and feed request item into worker item which is a Channel of request, in `engine/worker`, the worker function is going to fecth and parse the request. Please refer to the diagram below. 
  
    <img src="./Images/QueuedScheduler.png" height=60% width=60%>