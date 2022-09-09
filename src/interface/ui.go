package ui

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

func _error(err interface{}) {

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

}

func GetBanner() {
	
	content, err := ioutil.ReadFile("src/interface/banner.txt")
	_error(err)

	bannerMsg := color.New(color.FgCyan).Add(color.Bold)
	bannerMsg.Println(string(content))

}