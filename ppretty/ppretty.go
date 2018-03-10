package ppretty

import (
	"container/list"
	"fmt"
	"log"
	"strings"
)

const (
	AlignLeft = iota
	AlignCenter
	AlignRight
)

type PrintTable struct {
	ls *list.List
}

// Header add header to table
func (pt *PrintTable) Header(row []string) {
	pt.ls.PushFront(row)
}

func (pt *PrintTable) Append(row []string) {
	pt.ls.PushBack(row)
}

func (pt *PrintTable) Print(align int) {
	var tb [][]string
	for e := pt.ls.Front(); e != nil; e = e.Next() {
		val, ok := e.Value.([]string)
		if ok {
			tb = append(tb, val)
		}
		//pt.ls.Remove(e)
	}
	PrettyPrintAlign(tb, align)
}

func New() *PrintTable {
	return &PrintTable{list.New()}
}

// PrettyPrint print table
func PrettyPrint(tb [][]string) {
	PrettyPrintAlign(tb, AlignCenter)
}

func PrettyPrintAlign(tb [][]string, align int) {
	rows := len(tb)
	if rows == 0 {
		return
	}
	cols := len(tb[0])
	// 存储每列最长的长度
	colsWidth := make([]int, cols)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if colsWidth[i] <= strLen(tb[j][i]) {
				colsWidth[i] = strLen(tb[j][i]) + 2
			}
		}
	}
	printLine(colsWidth)
	for r := 0; r < rows; r++ {
		// 打印当前行
		printRow(tb[r], colsWidth, align)
		printLine(colsWidth)
	}
}

func printLine(colsWidth []int) {
	for i := 0; i < len(colsWidth); i++ {
		fmt.Print("+")
		for j := 0; j < colsWidth[i]; j++ {
			fmt.Print("-")
		}
	}
	fmt.Print("+\n")
}

func printRow(row []string, colsWidth []int, align int) {
	if len(row) != len(colsWidth) {
		log.Panic("parameter error")
	}
	for i := 0; i < len(row); i++ {
		fmt.Print("|")
		valLen := strLen(row[i])
		switch align {
		case AlignLeft:
			spaces := colsWidth[i] - valLen
			val := row[i] + strings.Repeat(" ", spaces)
			fmt.Print(val)
		case AlignCenter:
			left := (colsWidth[i] - valLen) / 2
			right := colsWidth[i] - valLen - left
			val := strings.Repeat(" ", left) + row[i] + strings.Repeat(" ", right)
			fmt.Print(val)
		case AlignRight:
			spaces := colsWidth[i] - valLen
			val := strings.Repeat(" ", spaces) + row[i]
			fmt.Print(val)
		default:
			log.Fatal("align parameter invalid")
			break
		}
	}
	fmt.Println("|")
}

func strLen(s string) int {
	l := 0
	for _, c := range s {
		if c < 0xFF {
			l += 1
		} else {
			l += 2
		}
	}
	return l
}
