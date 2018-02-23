package main

import (
	"fmt"

	"github.com/marlondu/gtool/cldmusic"
)

const (
	help = `usage: do it as tips say`
)

func main() {
	fmt.Println(help)
	//cldmusic.DoSearch()
	//s := time.Now()
	//cldmusic.VideoInfo("CF3F588B742A27FF4EADB208160904B8", cldmusic.VIDEO_TYPE)
	//fmt.Printf("spend time: %.3f/s\n", time.Since(s).Seconds())
	cldmusic.Help()
}
