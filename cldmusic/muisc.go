package cldmusic

import (
	"flag"
	"fmt"
	"log"
)

func DoSearch() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 3 {
		fmt.Println("at least three arguments")
		return
	}
	opt := args[0]
	tp := args[1]
	word := args[2]
	if opt == "search" {
		if tp == "song" {
			content, err := Search(word, SONG)
			if err != nil {
				log.Panic(err)
			}
			songs := Convert2Song(content)
			fmt.Println("---name---|----ID---|--artist--|---album---")
			for _, s := range songs {
				fmt.Println(s.String())
			}
		} else if tp == "video" {
			fmt.Println("---name---|----ID---|---type--|----artist--")
			content, err := Search(word, VIDEO)
			if err != nil {
				log.Panic(err)
			}
			videos := Convert2Video(content)
			for _, v := range videos {
				fmt.Println(v.String())
			}
		} else if tp == "mv" {
			fmt.Println("---name---|----ID---|--artist--")
			content, err := Search(word, MV_)
			if err != nil {
				log.Panic(err)
			}
			videos := Convert2MV(content)
			for _, v := range videos {
				fmt.Println(v.String())
			}
		}
	} else if opt == "download" {
		//id := word
		if tp == "video" {

		} else if tp == "mv" {

		}
	}
}
