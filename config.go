package main

import (
	"flag"
	"log"
	"os"
)

var (
	configFile        = ""
	defaultConfigFile = "/etc/jimput/jimput.toml"
	confdir           string
	config            Config
)

type Config struct {
	ConfDir string `toml:confdir`
}

func init() {
	flag.StringVar(&confdir, "confdir", "/etc/jimput", "jimput conf directory")
}

func initConfig() error {
	if configFile == "" {
		if _, err := os.Stat(defaultConfigFile); !os.IsNotExist(err) {
			configFile = defaultConfigFile
		}
	}

	// set defaults
	config = Config{
		ConfDir: "/etc/jimput",
	}

	// Update config from TOML configuration file
	if configFile == "" {
		log.Println("")
	}
	return nil
}
