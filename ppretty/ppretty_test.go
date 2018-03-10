package ppretty

import (
	"container/list"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	var tb = [][]string{
		[]string{"ID", "名称", "歌手", "类型"},
		[]string{"5451224", "Sunshine Girl", " moumoon", "MV"},
		[]string{"FFB3FF3FD143C528AF862662A4D72A1F", "火遍大街小巷的歌曲《Sunshine Girl》，妹子好萌", "音乐诊疗室", "video"},
	}
	PrettyPrint(tb)
}

func TestPrintTable_Print(t *testing.T) {
	var pt = &PrintTable{list.New()}
	pt.Append([]string{"5451224", "Sunshine Girl", " moumoon", "MV"})
	pt.Append([]string{"FFB3FF3FD143C528AF862662A4D72A1F", "火遍大街小巷的歌曲《Sunshine Girl》，妹子好萌", "音乐诊疗室", "video"})
	pt.Header([]string{"ID", "名称", "歌手", "类型"})
	pt.Print(AlignLeft)
	pt.Print(AlignRight)
}
