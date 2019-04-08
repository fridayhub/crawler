package fetcher

import (
	"bufio"
	"fmt"
	"log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)
func Fetcher(url string) ([]byte, error)  {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request error:%d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	//检测并转换编码
	utf8Reader := transCoding(bodyReader)

	return ioutil.ReadAll(utf8Reader)
}


func transCoding(rdata *bufio.Reader) io.Reader {
	//检测并转换编码
	e := determineEncoding(rdata)
	utf8Reader := transform.NewReader(rdata, e.NewDecoder())
	return utf8Reader
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fecther error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}