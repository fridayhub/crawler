package parser

import (
	"github.com/hakits/crawler/engine"
	"regexp"
)

var BaseUrl = "https://www.zhipin.com"
var cityListRe = regexp.MustCompile(`<a href="(/c[0-9]+/)" ka="sel-city-[0-9]+">([^<]+)</a>`)
var businessRe = regexp.MustCompile(`<a[ ]+href="(/c[0-9]+/b_[0-9A-Z%]+/)"[^ka]*ka="sel-business-[0-9]+">([^>]+)</a>`)
var JobListRe = regexp.MustCompile(`<div class="job-primary">
                                    <div class="info-primary">
                                        <h3 class="name">
                                            <a href="([^"]+)".*[\s ]+<div class="job-title">([^<]+)</div>`)

func ParseCityList(contents []byte) engine.ParseResult {
	match := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range match {
		//fmt.Printf("href:%s, city:%s\n", v[1], v[2])
		if string(v[2]) == "全国" {
			continue
		}
		result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        BaseUrl + string(v[1]),
			ParserFunc: ParseBusinessList,
		})
		break
	}
	return result
}

func ParseBusinessList(contents []byte) engine.ParseResult {
	match := businessRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range match {
		//fmt.Printf("href:%s, business:%s\n", v[1], v[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        BaseUrl + string(v[1]),
			ParserFunc: ParseJobList,
		})
		result.Items = append(result.Items, string(v[2]))
		break
	}
	return result
}

func ParseJobList(contents []byte) engine.ParseResult{
	result := engine.ParseResult{}

	match := JobListRe.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", match)
	for _, v := range match {
		//fmt.Printf("%s\n", v[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:BaseUrl + string(v[1]),
			ParserFunc:ParseProfile,
		})
		result.Items = append(result.Items, string(v[2]))
	}
	return result
}