package cldmusic

import (
	"bufio"
	"fmt"
	"os"
)

func Help() {
	const (
		stepOne = "what do you want to do (search/download/exit)? "
	)
	var param string
	reader := bufio.NewReader(os.Stdin)
exit:
	for {
		fmt.Print(stepOne)
		data, _, _ := reader.ReadLine()
		param = string(data)
		switch param {
		case "search":
			fmt.Print("what type do you want to search(video/mv/song/exit)? ")
			data, _, _ = reader.ReadLine()
			param = string(data)
			var tp int
			switch param {
			case "video":
				tp = VIDEO
			case "mv":
				tp = MV_
			case "song":
				tp = SONG
			case "exit":
				fmt.Println("exit ...")
				break exit
			default:
				continue
			}
			fmt.Print("please type in what you want to search: ")
			data, _, _ = reader.ReadLine()
			param = string(data)
			Search(param, tp)
		case "download":
			//fmt.Printf("please type in the id you want to download: ")
			fmt.Print("what do you want to download(video/mv)?")
			data, _, _ = reader.ReadLine()
			param = string(data)
			var tp int
			if param == "video" {
				tp = VIDEO_TYPE
			} else if param == "mv" {
				tp = MV_TYPE
			} else {
				continue
			}
			fmt.Print("please input the id you want to download:")
			data, _, _ = reader.ReadLine()
			param = string(data)
			if len(param) == 0 {
				fmt.Println("id can not be empty,please try it again.")
				continue
			}
			Download(param, tp)
		case "exit":
			fmt.Println("exit ...")
			break exit
		default:
			continue
		}
	}
}
