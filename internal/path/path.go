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

type Program struct {
	Id       int
	Name     string
	LastPath string
	LastTime string
}

func (p *Program) NewPath() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	p.Id = 1
	p.Name = path[strings.LastIndex(path, "/")+1:]
	p.LastPath = path
	p.LastTime = time.Now().Format("2006-01-02 15:04:05")
}

func SavePath() Program {
	p := Program{}
	p.NewPath()
	return p
}

func GetInfo() Program {
	file = os.Getenv("HOME") + "/." + file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var p Program
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
	fmt.Printf("When:\t\t%s\n", p.LastTime)
}

// print shot info path
func PrintShort() {
	p := GetInfo()
	fmt.Printf("%s\n", p.LastPath)
}

func GetMove() {
	p := GetInfo()
	fmt.Println(p.LastPath)
}

func FileSave() {
	p := SavePath()
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
