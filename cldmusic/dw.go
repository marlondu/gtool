package cldmusic

import (
	"fmt"
	"log"
	"net/http"

	"bufio"

	"os"

	"io"

	"github.com/marlondu/gtool/dson"
)

const (
	MV_TYPE = iota
	VIDEO_TYPE
)

type videoInfo struct {
	name       string
	id         string
	url        string // 下载地址
	resolution int    // 最大分辨率
	tp         int    // 类型
}

func VideoInfo(id string, tp int) map[string]string {
	info := FindVideoInfo(id, tp)
	uri := FindVideoUrl(id, info.resolution, tp)
	return map[string]string{
		"name": info.name,
		"url":  uri,
	}
}

func FindVideoInfo(id string, tp int) videoInfo {
	uri := "http://music.163.com/weapi/cloudvideo/v1/video/detail?csrf_token="
	param := fmt.Sprintf(`{"id":"%s","type":"%d","csrf_token":""}`, id, tp)
	if tp == MV_TYPE {
		uri = "http://music.163.com/weapi/v1/mv/detail?csrf_token="
		param = fmt.Sprintf(`{"id":"%s","type":"%d","csrf_token":""}`, id, tp)
	}

	resp, err := fetch(param, uri)
	if err != nil {
		log.Fatalln(err)
	}

	// "{"rid":"R_MV_5_5363527","offset":"0","total":"true","limit":"20","csrf_token":""}"
	// "{"id":"5363527","type":"0","csrf_token":""}"
	// "{"type":"MP4","id":"5363527","csrf_token":""}"
	// "{"id":"5363527","r":"480","csrf_token":""}"
	//fmt.Println(resp)
	if tp == MV_TYPE {
		return getMvInfo(resp)
	}
	return getVideoInfo(resp)
}

func getMvInfo(resp string) videoInfo {
	jsObj := dson.ParseObject(resp)
	data := jsObj.GetObject("data")
	r := 0
	rs := data.GetArray("brs")
	for i := 0; i < rs.Size(); i++ {
		robj := rs.GetObject(i)
		if r < robj.GetInt("br") {
			r = robj.GetInt("br")
		}
	}
	return videoInfo{
		id:         data.GetString("id"),
		name:       data.GetString("name"),
		tp:         1,
		resolution: r,
	}
}

func getVideoInfo(resp string) videoInfo {
	jsObj := dson.ParseObject(resp)
	data := jsObj.GetObject("data")
	r := 0
	rs := data.GetArray("resolutions")
	for i := 0; i < rs.Size(); i++ {
		robj := rs.GetObject(i)
		if r < robj.GetInt("resolution") {
			r = robj.GetInt("resolution")
		}
	}
	return videoInfo{
		id:         data.GetString("vid"),
		name:       data.GetString("title"),
		tp:         1,
		resolution: r,
	}
}

func FindVideoUrl(id string, resolution int, tp int) string {
	uri := "http://music.163.com/weapi/cloudvideo/playurl?csrf_token="
	param := fmt.Sprintf(`{"ids":"[\"%s\"]","resolution":"%d","csrf_token":""}`, id, resolution)
	if tp == MV_TYPE {
		uri = "http://music.163.com/weapi/song/enhance/play/mv/url?csrf_token="
		param = fmt.Sprintf(`{"id":"%s","r":"%d","csrf_token":""}`, id, resolution)
	}
	resp, err := fetch(param, uri)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(resp)
	jsObj := dson.ParseObject(resp)
	if tp == MV_TYPE {
		data := jsObj.GetObject("data")
		return data.GetString("url")
	}
	return jsObj.GetArray("urls").GetObject(0).GetString("url")
}

func Download(id string, tp int) {
	info := VideoInfo(id, tp)
	name := info["name"]
	vurl := info["url"]
	fmt.Printf("uri: %s, name: %s\n", vurl, name)
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
