package cldmusic

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marlondu/gtool/gsoup"
)

const (
	MV_TYPE = iota
	VIDEO_TYPE
)

func VideoInfo(id string, tp int) map[string]string {
	var vurl string
	if tp == 1 {
		vurl = fmt.Sprintf("http://music.163.com/video?id=%s", id)
	} else if tp == 0 {
		vurl = fmt.Sprintf("http://music.163.com/mv?id=%s", id)
	} else {
		log.Panic("tp must be 0 or 1")
	}
	resp, err := http.Get(vurl)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	node, err := gsoup.Parse(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	/*h2 := node.GetElementById("flag_title1")
	if h2 != nil {
		name := h2.Attribute("title")
		fmt.Println(name)
	}*/
	/*div := node.GetElementById("flash_box")
	if div != nil {
		hurl := div.Attribute("data-flashvars")
		// hurl, err = url.QueryUnescape(hurl)
		fmt.Println(hurl)
		arr := strings.Split(hurl, "&")
		for _, str := range arr {
			fmt.Println(str)
		}
	}*/

	nodes := node.Select("#flash_box")
	for _, n := range nodes {
		fmt.Println(n.Html())
	}
	/*doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	name, ext := doc.Find("#flag_title1").Attr("title")
	if ext {
		fmt.Println(name)
	}
	hurl, ext := doc.Find("#flash_box").Attr("data-flashvars")
	if ext {
		hurl, err = url.QueryUnescape(hurl)
		fmt.Println(hurl)
	}*/

	return nil
}
