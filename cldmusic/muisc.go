package cldmusic

import (
	"bufio"
	"fmt"
	"os"
)

func Help() {
	const (
		stepOne = "what do you want to do (search|s/download|d/exit|e)? "
	)
	var param string
	reader := bufio.NewReader(os.Stdin)
exit:
	for {
		fmt.Print(stepOne)
		data, _, _ := reader.ReadLine()
		param = string(data)
		switch param {
		case "search", "s":
			fmt.Print("what type do you want to search(video|v/mv|v/song|s/exit|e)? ")
			data, _, _ = reader.ReadLine()
			param = string(data)
			var tp int
			switch param {
			case "video", "v":
				tp = VIDEO
			case "mv", "m":
				tp = MV_
			case "song", "s":
				tp = SONG
			case "exit", "e":
				fmt.Println("exit ...")
				break exit
			default:
				continue
			}
			fmt.Print("please type in what you want to search: ")
			data, _, _ = reader.ReadLine()
			param = string(data)
			Search(param, tp)
		case "download", "d":
			//fmt.Printf("please type in the id you want to download: ")
			fmt.Print("what do you want to download(video|v/mv|m)?")
			data, _, _ = reader.ReadLine()
			param = string(data)
			var tp int
			if param == "video" || param == "v" {
				tp = VIDEO_TYPE
			} else if param == "mv" || param == "m" {
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
		case "exit", "e":
			fmt.Println("exit ...")
			break exit
		default:
			continue
		}
	}
}
