package main

import (
	"lastpath/internal/path"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		path.ChangeCurrentPath()
		return
	}

	switch os.Args[1] {
	case "save":
		path.SavePath()
	case "info":
		path.PrintInfo()
	case "move":
		path.ChangeCurrentPath()
	}
}
