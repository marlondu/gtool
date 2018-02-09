package dson

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestParseObject(t *testing.T) {
	url := "https://m.douban.com/rexxar/api/v2/gallery/subject_feed?start=0&count=4&subject_id=26942674&ck=null"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Cookie", `ll="108288"; bid=0BPOu-pO5Gk; __utma=30149280.1578426299.1516892582.1516892582.1516976747.2; __utmc=30149280; __utmz=30149280.1516976747.2.2.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmt=1; __utmt_douban=1; __utmt_t1=1; ct=y; viewed="1788421_5333562_1148282_27601596"; __utmb=30149280.18.10.1516976747`)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	req.Header.Add("Referer", "https://movie.douban.com/subject/26942674/?from=showing")
	req.Header.Add("Host", "m.douban.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Origin", "https://movie.douban.com")
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	jsonStr := string(data)

	resultJson := ParseObject(jsonStr)
	items := resultJson.GetArray("items")
	for i := 0; i < items.Size(); i++ {
		item := items.GetObject(i)
		topic := item.GetObject("topic")
		name := topic.GetString("name")
		t.Log(name)
	}
}
