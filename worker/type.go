package worker

import (
	"fmt"
	"github.com/hakits/crawler/config"
	"github.com/hakits/crawler/engine"
	"github.com/hakits/crawler/zhipin/parser"
	"github.com/kataras/iris/core/errors"
	"log"
	"reflect"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url:r.Url,
		Parser:SerializedParser{
			Name:name,
			Args:args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items:r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializedRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:r.Url,
		Parser:parser,
	}, nil
}

func DeserializeResult(r ParseResult) (engine.ParseResult) {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("err deserializeing reqeust: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseBusinessList:
		return engine.NewFuncParser(parser.ParseBusinessList, config.ParseBusinessList), nil
	case config.ParseJobList:
		return engine.NewFuncParser(parser.ParseJobList, config.ParseJobList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ProfileParser:
		log.Printf("type:%v", reflect.TypeOf(p.Args))
		if args, ok := p.Args.([]interface{}); ok {
			arg := []string{}
			for _, v := range args {
				arg =append(arg, v.(string))
			}
			return parser.NewProfileParser(arg), nil
		} else {
			return nil, fmt.Errorf("invaild args:%v", p.Args)
		}
	default:
		return nil, errors.New("unkown paser name")

	}
}

