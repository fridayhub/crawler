package engine

type Request struct {
	Url        string
	Parser Parser
}

type ParserFunc func(content []byte) ParseResult

type Parser interface {
	Parse(content []byte) ParseResult
	Serialize() (name string, args interface{})
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url string
	Id string
	Payload interface{}
}

type NilParser struct {}

func (NilParser) Parse(_ []byte) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "nilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name string
}

func (f *FuncParser) Parse(content []byte) ParseResult {
	return  f.parser(content)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser:p,
		name:name,
	}
}
