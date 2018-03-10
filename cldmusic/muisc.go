package cldmusic

import (
	"bufio"
	"fmt"
	"os"
)

func Help() {
	const (
		stepOne = "请输入你要进行的操作(search|s/download|d/exit|e)? "
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
			fmt.Print("请输入要搜索的类型(video|v/mv|v/song|s/exit|e)? ")
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
			fmt.Print("请输入要搜索的关键字: ")
			data, _, _ = reader.ReadLine()
			param = string(data)
			Search(param, tp)
		case "download", "d":
			//fmt.Printf("please type in the id you want to download: ")
			fmt.Print("请输入要下载的类型(video|v/mv|m)?")
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
			fmt.Print("请输入要下载的ID:")
			data, _, _ = reader.ReadLine()
			if data == nil || len(data) == 0 {
				fmt.Println("ID不能为空，请重试.")
				continue
			}
			param = string(data)
			Download(param, tp)
		case "exit", "e":
			fmt.Println("exit ...")
			break exit
		default:
			continue
		}
	}
}
