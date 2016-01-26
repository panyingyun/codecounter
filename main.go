package main

import (
	"fmt"

	"github.com/panyingyun/codecounter/fsch"
)

func main() {
	fmt.Println("Code counter here!!!")
	fs := fsch.NewFileSearch()
	fs.Dir = "E:\\FocusVIew\\trunk\\src\\com\\lyk\\focusview_master"
	fs.KeyWord = "*.java"
	fl, _ := fs.SearchToList()
	for _, fn := range fl {
		fmt.Println(fn)
	}
}
