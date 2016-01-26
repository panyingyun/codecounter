package main

import (
	"fmt"

	"github.com/panyingyun/codecounter/fsch"
)

func main() {
	//统计后缀
	suffixs := []string{"*.c", "*.cpp", "*.h", "*.go", "*.java", "*.cc"}

	fs := fsch.NewFileSearch()
	fs.Dir = "E:\\mygithub\\um\\trunk"
	fs.SubDir = true
	fs.CaseMind = false

	for _, suffix := range suffixs {
		fs.KeyWord = suffix
		fl, _ := fs.SearchToList()
		for _, fn := range fl {
			fmt.Println(fn)
		}
	}
}
