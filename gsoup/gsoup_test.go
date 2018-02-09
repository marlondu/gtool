package gsoup

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	var vurl = "http://music.163.com/video?id=CF3F588B742A27FF4EADB208160904B8"

	resp, err := http.Get(vurl)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	node, err := Parse(resp.Body)
	if err != nil {
		t.Error(err)
	}

	nodes := node.Select("#flash_box")
	//fmt.Println(len(nodes))
	for _, n := range nodes {
		t.Log(n.Html())
	}

	div := node.GetElementById("flash_box")
	if div != nil {
		hurl := div.Attribute("data-flashvars")
		// hurl, err = url.QueryUnescape(hurl)
		fmt.Println(hurl)
		arr := strings.Split(hurl, "&")
		for _, str := range arr {
			fmt.Println(str)
		}
	}
}
