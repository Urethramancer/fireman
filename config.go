package main

import (
	"fmt"
	"os"

	"code.google.com/p/gcfg"
)

const BASE string = "fireman"
const CONFIGNAME string = BASE + ".conf"

type Config struct {
	Main struct {
		Server string
		Key    string
	}
}

var cfg Config
var cfgfile *string

func (conf *Config) Save() error {
	f, err := os.Create(*cfgfile)
	if err != nil {
		return err
	}
	f.WriteString("[main]\n")
	fmt.Fprintf(f, "server = %s\n", cfg.Main.Server)
	fmt.Fprintf(f, "key = %s\n", cfg.Main.Key)

	f.Close()
	os.Chmod(*cfgfile, 0700)
	return nil
}

func loadConfig(name *string) {
	cfg.Main.Server = "http://localhost:9090/plugins/userService/"
	cfg.Main.Key = "<key>"

	reloadConfig(name)
}

func reloadConfig(name *string) {
	if !exists(*name) {
		p("Configuration file %s does not exist. Generating a default file.", *name)
		p("Creating default %s", *name)
		cfg.Save()
	}

	err := gcfg.ReadFileInto(&cfg, *name)
	if err != nil {
		f("Configuration error: %s", err.Error())
	}
}
