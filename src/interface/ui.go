package ui

import (
	"github.com/fatih/color"
)

func GetBanner() {

	text := `

    \|/ ____ \|/       
     @~/ ,. \~@     pwnfaces v1.3.7a - Primefaces 5.X EL Injection Exploit
    /_( \__/ )_\    by oppsec & march0s1as (thanks to pimps)
       \__U_/     
	
	`

	bannerMsg := color.New(color.FgCyan).Add(color.Bold)
	bannerMsg.Println(text)

}
