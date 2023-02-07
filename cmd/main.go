package main

import (
	"fmt"
	"lastpath/internal/path"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("move")
		path.ChangeCurrentPath()
		return
	}

	switch os.Args[1] {
	case "save":
		fmt.Println("save")
		path.SavePath()
	case "info":
		fmt.Println("info")
		path.GetInfo()
	case "move":
		fmt.Println("move")
		path.ChangeCurrentPath()
	}
}
