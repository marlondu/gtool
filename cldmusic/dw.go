package cldmusic

import (
	"fmt"
	"log"
	"net/http"

	"strings"

	"bufio"

	"os"

	"io"

	"github.com/marlondu/gtool/gsoup"
	"qiniupkg.com/x/url.v7"
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
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != 200 {
		log.Panic(err)
	}
	node, err := gsoup.Parse(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	div := node.GetElementById("flash_box")
	info := div.Attribute("data-flashvars")
	infoArr := strings.Split(info, "&")
	var uri string
	var name string
	for _, ia := range infoArr {
		if strings.HasPrefix(ia, "hurl") {
			uri = strings.Replace(ia, "hurl=", "", -1)
		}
		if strings.HasPrefix(ia, "murl") && len(uri) == 0 {
			uri = strings.Replace(ia, "murl=", "", -1)
		}
		if strings.HasPrefix(ia, "trackName") {
			name = strings.Replace(ia, "trackName=", "", -1)
		}
	}
	uri, _ = url.Unescape(uri)
	return map[string]string{
		"name": name,
		"url":  uri,
	}
}

func Download(id string, tp int) {
	info := VideoInfo(id, tp)
	name := info["name"]
	vurl := info["url"]
	//fmt.Printf("uri: %s, name: %s\n", vurl, name)
	savePath := fmt.Sprintf("E:/video/mv/%s.mp4", name)
	resp, err := http.Get(vurl)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	file, _ := os.Create(savePath)
	writer := bufio.NewWriter(file)
	var buff = make([]byte, 4096)
	var size = resp.ContentLength
	var per float64 // 下载百分比
	var bytes int
	for {
		if n, err := reader.Read(buff); err != io.EOF {
			writer.Write(buff[:n])
			bytes += n
			per = float64(bytes) / float64(size) * 100
			fmt.Printf("正在下载: %s / %.2f%%\r", savePath, per)
		} else {
			fmt.Printf("%s 下载完毕\n", savePath)
			writer.Flush()
			file.Close()
			break
		}
	}
}
