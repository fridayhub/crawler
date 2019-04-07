package parser

import (
	"github.com/hakits/crawler/engine"
	"github.com/hakits/crawler/model"
	"regexp"
)

var JobNameRe = regexp.MustCompile(`<title>([^<]+)</title>`) //match[0][1]
//CompanyRe := `<a ka="job-detail-company".*[\s].*target="_blank">[\s ]+(.*)[\s ]+</a>` //match[1][1]
var CompanyRe = regexp.MustCompile(`company:'(.*)',`)                     //
var ScaleRe = regexp.MustCompile(`<p><i class="icon-scale"></i>(.*)</p>`) //match[0][1]
var SalaryRe = regexp.MustCompile(`job_salary: '([0-9K-]+)'`)             //match[0][1]

var LoYeEdRe = regexp.MustCompile(` <p>(.*)<em class="dolt"></em>(.*)<em class="dolt"></em>(.*)</p>`)
var JobTagsRe = regexp.MustCompile(`<div class="job-tags">[\s ]+(.*)[\s ]+</div>`) //match[0][1]
var JobSecRe = regexp.MustCompile(`<div class="text">[\s ]+(.*)[\s ]+</div>`)
var RecruiterRe = regexp.MustCompile(` </div>
                    <h2 class="name">(.*)<i class="icon-vip"></i></h2>
                    <p class="gray">(.*)<em class="vdot">·</em>.*</p>
                </div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	profile.JobName = string(regxItem(JobNameRe, contents)[0][1])
	profile.Company = string(regxItem(CompanyRe, contents)[0][1])
	profile.Scale = string(regxItem(ScaleRe, contents)[0][1])
	profile.JobName = string(regxItem(JobNameRe, contents)[0][1])
	profile.Salary = string(regxItem(SalaryRe, contents)[0][1])
	tmpLYE := regxItem(LoYeEdRe, contents)
	profile.Location = string(tmpLYE[0][1])  //地点
	profile.Years = string(tmpLYE[0][3])     //工作年限
	profile.Education = string(tmpLYE[0][3]) //教育程度
	profile.JobTags = string(regxItem(JobTagsRe, contents)[0][1])
	profile.JobSec = string(regxItem(JobSecRe, contents)[0][1])
	tmpRec := regxItem(RecruiterRe, contents)
	profile.Recruiter = string(tmpRec[0][1]) + "|" + string(tmpRec[0][2])

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func regxItem(re *regexp.Regexp, contents []byte) [][][]byte {
	match := re.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", match)
	if len(match) >= 1 {
		return match
	} else {
		result := make([][][]byte, 0, 3)
		slice := make([][]byte, 3)
		for j := range slice {
			slice[j] = []byte{}
		}
		result = append(result, slice)

		return result
	}
}
