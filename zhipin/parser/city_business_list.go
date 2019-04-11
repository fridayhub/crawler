package parser

import (
	"github.com/hakits/crawler/engine"
	"regexp"
	"strconv"
)

var BaseUrl = "https://www.zhipin.com"
var cityListRe = regexp.MustCompile(`<a href="(/c[0-9]+/)" ka="sel-city-[0-9]+">([^<]+)</a>`)
var businessRe = regexp.MustCompile(`<a[ ]+href="(/c[0-9]+/b_[0-9A-Z%]+/)"[^ka]*ka="sel-business-[0-9]+">([^>]+)</a>`)
var JobListRe = regexp.MustCompile(`<div class="job-primary">
                                    <div class="info-primary">
                                        <h3 class="name">
                                            <a href="([^"]+)".*[\s ]+<div class="job-title">([^<]+)</div>`)

const query = "golang"

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
		break //For testing,just get one city
	}
	return result
}

func ParseBusinessList(contents []byte) engine.ParseResult {
	match := businessRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range match {
		//fmt.Printf("href:%s, business:%s\n", v[1], v[2])
		//fmt.Printf("BusinessList:%s\n", BaseUrl + string(v[1]))

		//Take 10 page for each business
		for i := 1; i <= 10; i++ {
			result.Requests = append(result.Requests, engine.Request{
				Url:        BaseUrl + string(v[1]) + "?page=" + strconv.Itoa(i) + "&query=" + query,
				ParserFunc: ParseJobList,
			})
			result.Items = append(result.Items, string(v[2]))
		}
		break //For testing, just get one business
	}
	return result
}

func ParseJobList(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	match := JobListRe.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", match)
	for _, v := range match {
		//fmt.Printf("%s\n", v[1])
		jobName := string(v[2])
		uri := string(v[1])
		url := BaseUrl + uri
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(content []byte) engine.ParseResult { //只是把函数赋值给ParseFunc 现在并不运行,engine调度的时候才运行
				return ParseProfile(content, jobName, url)
			},
		})
		result.Items = append(result.Items, jobName)
	}
	return result
}
