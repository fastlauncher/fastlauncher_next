package main

import (
	"flag"

	"github.com/probeldev/fastlauncher/log"
	"github.com/probeldev/fastlauncher/mode"
	"github.com/probeldev/fastlauncher_next/ui"
)

func main() {
	cfgPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if cfgPath != nil && *cfgPath != "" {
		ca := mode.ConfigMode{}
		apps := ca.GetFromFile(*cfgPath)
		ui.StartUI(apps)
		return
	}

	oa := mode.OsMode{}
	apps, err := oa.GetAll()
	if err != nil {
		// TODO
		log.Println(err)
	}

	ui.StartUI(apps)
}
