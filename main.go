package main

import (
	"fmt"
	"github.com/hakits/crawler/engine"
	"github.com/hakits/crawler/zhipin/parser"
	"regexp"
)

func main() {

	engine.Run(engine.Request{
		Url:        parser.BaseUrl + "/c101010100/",
		ParserFunc: parser.ParseCityList,
	})


		//body, err := fetcher.Fetcher("https://www.zhipin.com/c101010100/b_%E6%9C%9D%E9%98%B3%E5%8C%BA/")
		//if err != nil {
		//	log.Printf("Fether:error fetch url %v", err)
		//}
		//printAreaList(body)

}

func printAreaList(contents []byte) {
	//fmt.Printf("%s",contents)
	//re := regexp.MustCompile(`<a[ ]+href="(/c101010100/a_[0-9A-Z%]+/)"[^ka]*ka="sel-area-[0-9]+">([^>]+)</a>`)
	//JobNameRe := `<title>([^<]+)</title>` //match[0][1]
	////CompanyRe := `<a ka="job-detail-company".*[\s].*target="_blank">[\s ]+(.*)[\s ]+</a>` //match[1][1]
	//CompanyRe := `company:'(.*)',` //
	//Scale := `<p><i class="icon-scale"></i>(.*)</p>` //match[0][1]
	//JobName := `<h1>(.*)</h1>` //match[0][1]
	//Salary := `job_salary: '([0-9K-]+)'` //match[0][1]
	//Location string //地点 match[0][1]
	//Years string //工作年限 match[0][2]
	//Education string  //教育程度 match[0][3]
	//LoYeEd := ` <p>(.*)<em class="dolt"></em>(.*)<em class="dolt"></em>(.*)</p>`
	//JobTags := `<div class="job-tags">[\s ]+(.*)[\s ]+</div>` //match[0][1]
	//JobSec := `<div class="text">[\s ]+(.*)[\s ]+</div>`
	//Recruiter := ` </div>
    //                <h2 class="name">(.*)<i class="icon-vip"></i></h2>
    //                <p class="gray">(.*)<em class="vdot">·</em>.*</p>
    //            </div>`

    JobListRe := `<div class="job-primary">
                                    <div class="info-primary">
                                        <h3 class="name">
                                            <a href="([^"]+)".*[\s ]+<div class="job-title">([^<]+)</div>`
	re := regexp.MustCompile(JobListRe)
	match := re.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", match)
	for _, v := range match {
		fmt.Printf("%s, %s\n", v[1], v[2])
	}
}
