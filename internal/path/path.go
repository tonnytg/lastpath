package path

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	file = "filepath.txt"
)

type Path struct {
	Id       int
	Name     string
	LastPath string
	CreateAt string
}

func (p *Path) NewPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	p.Id = 1
	p.Name = path[strings.LastIndex(path, "/")+1:]
	p.LastPath = path
	p.CreateAt = time.Now().Format("2006-01-02 15:04:05")
}

func SavePath() Path {
	p := Path{}
	p.NewPath()
	FileSave(p)
	return p
}

func GetInfo() Path {
	file = os.Getenv("HOME") + "/." + file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var p Path
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func PrintInfo() {
	p := GetInfo()

	fmt.Printf("ID:\t\t%d\n", p.Id)
	fmt.Printf("Name:\t\t%s\n", p.Name)
	fmt.Printf("LastPath:\t%s\n", p.LastPath)
	fmt.Printf("When:\t\t%s\n", p.CreateAt)
}

func PrintShort() {
	p := GetInfo()
	fmt.Printf("%s\n", p.LastPath)
}

func ChangeCurrentPath() {
	p := GetInfo()
	fmt.Println(p.LastPath)
}

func FileSave(p Path) {
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	file = os.Getenv("HOME") + "/." + file
	err = os.WriteFile(file, j, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
