package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KEVISONG/hello-go-web/pkg/api/debug"
	"github.com/KEVISONG/hello-go-web/pkg/api/server"
	"github.com/KEVISONG/hello-go-web/pkg/config"
)

var version = "1.0.0"

func main() {

	fConfigFile := flag.String(
		"config",
		"./config.yml",
		"path to config file, e.g. --config=./config.yml, default is ./config.yml",
	)
	fVersion := flag.Bool(
		"version",
		false,
		"prints version and exits",
	)

	flag.Parse()

	if *fVersion {
		fmt.Printf("Version %s\n", version)
		os.Exit(0)
	}

	if fConfigFile != nil {
		err := config.InitConfig(*fConfigFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	go server.Run(config.C.Server)

	debug.Run(config.C.Debug)

}
