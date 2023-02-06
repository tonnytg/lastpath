package main

import (
	"fmt"
	"os"

	"lastpath/internal/database"
	"lastpath/internal/entity"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("vazio")
		return
	}

	switch os.Args[1] {
	case "save":
		fmt.Println("save")
		path := entity.Path{
			ID:       1,
			Path:     "/home/tonnytg/teste1",
			CreateAt: "Agora",
		}

		db := database.NewDB()
		db.Create(&path)
	case "info":
		fmt.Println("info")
	case "move":
		fmt.Println("move")
	}
}
