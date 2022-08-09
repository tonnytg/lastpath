package cmd

import (
	"flag"
	"fmt"
	"lstp/internal/path"
)

func main() {
	fmt.Println("vim-go")
}

func CheckFlag() error {
	save := flag.Bool("save", false, "save current path")
	info := flag.Bool("info", false, "get last path info")
	move := flag.Bool("move", false, "path to move")
	flag.Parse()
	if *save {
		path.FileSave()
	}
	if *info {
		path.PrintInfo()
	}
	if *move {
		path.GetMove()
	}
	return nil
}
