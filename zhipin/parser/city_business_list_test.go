package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, e := ioutil.ReadFile("cylist_test_data.html")
	if e != nil {
		panic(e)
	}

	result := ParseCityList(contents)
	const resultSize = 14
	if len(result.Requests) != resultSize {
		t.Errorf("result shoud have %d requests; but had %d", resultSize, len(result.Requests))
	}
}
