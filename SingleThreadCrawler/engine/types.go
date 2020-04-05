package engine

// Parser interface that can serialize Parser into data that can
// be transmitted over network
type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// Type representing function that parses html contents
type ParserFunc func(
	contents []byte, url string) ParseResult

// Struct of to construct parsing function
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

type Request struct {
	Url    string
	Parser Parser
}

type Item struct {
	Payload interface{}
	Id      string
	Url     string
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (_ string, _ interface{}) {
	return "NilParser", nil
}
