package ui

import (
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

	text := `

    \|/ ____ \|/       
     @~/ ,. \~@     pwnfaces v1.3.2 - Primefaces 5.X EL Injection Exploit
    /_( \__/ )_\    by oppsec (thanks to pimps)
       \__U_/     
	
	`

	bannerMsg := color.New(color.FgCyan).Add(color.Bold)
	bannerMsg.Println(text)

}
