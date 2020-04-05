package worker

import (
	"GoLangIntro/DistributedCrawler/config"
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/parser"
	"errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

// Here, we need to re-define Request and ParseRequest so that the data can be
// transmitted through network
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// Serialize Request into struct above
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

// Serialize ParseResult into struct above
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

// Deserialize Request into engine.Request
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, nil
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCarDetail:
		return engine.NewFuncParser(
			parser.ParseCarDetail, config.ParseCarDetail), nil
	case config.ParseCarList:
		return engine.NewFuncParser(
			parser.ParseCarList, config.ParseCarList), nil
	case config.ParseCarModel:
		return engine.NewFuncParser(
			parser.ParseCarModel, config.ParseCarModel), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("Unknown Parser name")
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("Error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}
