package parser

import (
	"github.com/hakits/crawler/config"
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
		//result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        BaseUrl + string(v[1]),
			Parser: engine.NewFuncParser(ParseBusinessList, config.ParseBusinessList), //之前的parseFunc函数，改成Parser接口
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
				Parser: engine.NewFuncParser(ParseJobList, config.ParseJobList),
			})
			//result.Items = append(result.Items, string(v[2]))
		}
		break //For testing, just get one business
	}
	return result
}

type ProfileParser struct {
	jobName string
	url string
}

func (p *ProfileParser) Parse(content []byte) engine.ParseResult {
	return parseProfile(content, p.jobName, p.url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	var tmpargs []string
	tmpargs = append(tmpargs, p.jobName)
	tmpargs = append(tmpargs, p.url)
	//fmt.Printf("args:%v", tmpargs)
	return config.ProfileParser, tmpargs
}

func NewProfileParser(args []string) *ProfileParser  {
	return &ProfileParser{
		jobName:args[0],
		url: args[1],
	}
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
		var args []string
		args = append(args,jobName)
		args = append(args, url)

		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			Parser: NewProfileParser(args),
		})
		//result.Items = append(result.Items, jobName)
	}
	return result
}
