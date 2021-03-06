package cldmusic

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"strings"

	"strconv"

	"github.com/marlondu/gtool/dson"
	"github.com/marlondu/gtool/ppretty"
)

/**
// uriLog := "http://music.163.com/weapi/feedback/weblog?csrf_token="
// uriSug := "http://music.163.com/weapi/search/suggest/multimatch?csrf_token="
// uriSeach := "http://music.163.com/weapi/cloudsearch/get/web?csrf_token="
//keyLog := `{"logs":"[{\"action\":\"searchkeywordclient\",\"json\":{\"type\":\"song\",\"keyword\":\"恋曲1990\",\"offset\":0}}]","csrf_token":""}`
//keySug := `{"s":"恋曲1990","csrf_token":""}`
//keySearch := `{"hlpretag":"<span class=\"s-fc7\">","hlposttag":"</span>","s":"恋曲1990","type":"1","offset":"0","total":"true","limit":"30","csrf_token":""}`
//              {"hlpretag":"<span class=\"s-fc7\">","hlposttag":"</span>","s":"恋曲1990","type":"1014","offset":"0","total":"true","limit":"20","csrf_token":""}
case 1 : "song";
case 100: "artist";
case 10: "album";
case 1004: "mv";
case 1014: "video";
case 1006: "lyric";
case 1e3: "list";
case 1009: "djradio";
case 1002: "user";
*/

const (
	SONG    = 1
	ARTIST  = 100
	ALBUM   = 10
	MV_     = 1004
	VIDEO   = 1014
	LYRIC   = 1006
	LIST    = 1e3
	DJRADIO = 1009
	USER    = 1002
	//urlLog  = "http://music.163.com/weapi/feedback/weblog?csrf_token="
	//urlSug  = "http://music.163.com/weapi/search/suggest/multimatch?csrf_token="
	urlSea = "http://music.163.com/weapi/cloudsearch/get/web?csrf_token="
)

type Song struct {
	Name    string   `json:"name"`
	Id      int      `json:"id"`
	Artists []string `json:"artists"`
	Album   string   `json:"album"`
}

func (s *Song) String() string {
	return fmt.Sprintf(" [%s]  [%d] [%s] [ %s]", s.Name, s.Id, strings.Join(s.Artists, ","), s.Album)
}

type MV struct {
	Name    string   `json:"name"`
	Id      int      `json:"id"`
	Artists []string `json:"artists"`
}

func (v *MV) String() string {
	return fmt.Sprintf(" [%s]  [%d] [%s]", v.Name, v.Id, strings.Join(v.Artists, ","))
}

// Video is video include video and mv
// Type 0 is mv
// Type 1 is video
type Video struct {
	Title   string   `json:"title"`
	Id      string   `json:"id"`
	Artists []string `json:"artists"`
	Type    string   `json:"type"`
}

func (v *Video) String() string {
	return fmt.Sprintf(" [%s] [%s] [%s] [%s]", v.Title, v.Id, v.Type, strings.Join(v.Artists, ","))
}

func Convert2Song(content string) []Song {
	resJson := dson.ParseObject(content).GetObject("result")
	songs := resJson.GetArray("songs")
	results := make([]Song, songs.Size(), songs.Size())
	for i := 0; i < songs.Size(); i++ {
		s := songs.GetObject(i)
		name := s.GetString("name")
		id := s.GetInt("id")
		artists := s.GetArray("ar")
		var ars = make([]string, 0, 2)
		for j := 0; j < artists.Size(); j++ {
			ar := artists.GetObject(j)
			ars = append(ars, ar.GetString("name"))
		}
		al := s.GetObject("al")
		alName := al.GetString("name")
		song := Song{
			Name:    name,
			Id:      id,
			Artists: ars,
			Album:   alName,
		}
		results[i] = song
	}
	return results
}

func Convert2Video(content string) []Video {
	resJson := dson.ParseObject(content).GetObject("result")
	videos := resJson.GetArray("videos")
	results := make([]Video, videos.Size(), videos.Size())
	for i := 0; i < videos.Size(); i++ {
		v := videos.GetObject(i)
		title := v.GetString("title")
		id := v.GetString("vid")
		artists := v.GetArray("creator")
		tp := v.GetInt("type")
		tpVal := "mv"
		if tp == VIDEO_TYPE {
			tpVal = "video"
		}
		var ars = make([]string, 0, 2)
		for j := 0; j < artists.Size(); j++ {
			ar := artists.GetObject(j)
			ars = append(ars, ar.GetString("userName"))
		}
		video := Video{
			Title:   title,
			Id:      id,
			Artists: ars,
			Type:    tpVal,
		}
		results[i] = video
	}
	return results
}

func Convert2MV(content string) []MV {
	resJson := dson.ParseObject(content).GetObject("result")
	videos := resJson.GetArray("mvs")
	results := make([]MV, videos.Size(), videos.Size())
	for i := 0; i < videos.Size(); i++ {
		v := videos.GetObject(i)
		title := v.GetString("name")
		id := v.GetInt("id")
		artists := v.GetArray("artists")
		var ars = make([]string, 0, 2)
		for j := 0; j < artists.Size(); j++ {
			ar := artists.GetObject(j)
			ars = append(ars, ar.GetString("name"))
		}
		video := MV{
			Name:    title,
			Id:      id,
			Artists: ars,
		}
		results[i] = video
	}
	return results
}

func Search(word string, tp int) {
	param :=
		`{"hlpretag":"<span class=\"s-fc7\">","hlposttag":"</span>","s":"%s","type":"%d","offset":"0","total":"true","limit":"30","csrf_token":""}`
	param = fmt.Sprintf(param, word, tp)
	content, err := fetch(param, urlSea)
	if err != nil {
		log.Panic(err)
	}
	switch tp {
	case VIDEO:
		videos := Convert2Video(content)
		prettyPrintVideo(videos)
	case MV_:
		mvs := Convert2MV(content)
		prettyPrintMV(mvs)
	case SONG:
		songs := Convert2Song(content)
		prettyPrintSong(songs)
	default:
		fmt.Println("type value error")
		fmt.Println(content)
	}
	return
}

func fetch(param string, uri string) (string, error) {
	paramsVal := cloudEnc(param)
	params := url.Values{
		"params":    []string{paramsVal},
		"encSecKey": []string{rsaKey()},
	}
	/*proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}

	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}*/
	client := http.DefaultClient
	req, err := http.NewRequest("POST", uri, bytes.NewBufferString(params.Encode()))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "http://music.163.com/mv?id=5363527")
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return string(data), err
}

func rsaKey() string {
	return "95d58bb933dd1768c93cd5948b4dd64b8788e77ff9a6fe14cd059979b87a5085fdb9b02a4e0fa084b30a0cae6f3fd73a955b4edc26d55c74641e713f24255a4750d3715f42066e908c5898de6522079267e273895369eae0c39b96bae295cc32fe1acd309bf7936a20922e618aa9f49c4002d03e833b653bc69f4923c89701b7"
}

func cloudEnc(src string) string {
	key := []byte("0CoJUm6Qyw8W8jud")
	secKey := "XTiAA52B4hlTMdus"
	enc0 := aesEncrypt(key, []byte(src))
	return aesEncrypt([]byte(secKey), []byte(enc0))
}

func aesEncrypt(key []byte, src []byte) string {
	iv := []byte("0102030405060708") // 0102030405060708
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}
	// var src = []byte(`{"csrf_token":""}`)
	src = PKCS5Padding(src, block.BlockSize())

	encrypter := cipher.NewCBCEncrypter(block, iv)
	var dst = make([]byte, len(src))
	encrypter.CryptBlocks(dst, src)
	encoded := base64.StdEncoding.EncodeToString(dst)
	return encoded
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	paddingSize := blockSize - len(src)%blockSize
	paddingBytes := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(src, paddingBytes...)
}

func UnPKCS5Padding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

func prettyPrintMV(list []MV) {
	// 每列的最大长度
	var tb [][]string
	head := []string{"ID", "名称", "歌手"}
	tb = append(tb, head)
	for _, v := range list {
		row := []string{strconv.Itoa(v.Id), v.Name, strings.Join(v.Artists, ",")}
		tb = append(tb, row)
	}
	ppretty.PrettyPrint(tb)
}

func prettyPrintVideo(list []Video) {
	var tb [][]string
	head := []string{"ID", "名称", "歌手", "类型"}
	tb = append(tb, head)
	for _, v := range list {
		row := []string{v.Id, v.Title, strings.Join(v.Artists, ","), v.Type}
		tb = append(tb, row)
	}
	ppretty.PrettyPrint(tb)
}

func prettyPrintSong(list []Song) {
	var tb [][]string
	head := []string{"ID", "名称", "歌手", "专辑"}
	tb = append(tb, head)
	for _, s := range list {
		var row = []string{strconv.Itoa(s.Id), s.Name, strings.Join(s.Artists, ","), s.Album}
		tb = append(tb, row)
	}
	ppretty.PrettyPrint(tb)
}
