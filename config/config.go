// Package config is responsible for jimput configuration, either from the command line or
// from configuration file.
//
// Thanks to https://github.com/kelseyhightower/confd for the inspiration (maybe that's an
// understatement).
package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/tdeckers/jimput/log"
	"io/ioutil"
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
	Debug   bool   `toml:debug`
	Quiet   bool   `toml:quiet`
	Verbose bool   `toml:verbose`
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
		log.Warning("Skipping config file.")
	} else {
		log.Debug("Loading " + configFile)
		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
		_, err = toml.Decode(string(configBytes), &config)
		if err != nil {
			return err
		}
	}

	// Update config from command line flags
	processFlags()

	// Configure logging.
	log.SetQuiet(config.Quiet)
	log.SetVerbose(config.Verbose)
	log.SetDebug(config.Debug)

	// Template configuration
	// Last todo

	return nil
}

// processFlags iterates through each flag set on the command line and
// overrides corresponding configuration settings.
func processFlags() {
	flag.Visit(setConfigFromFlag)
}

func setConfigFromFlag(f *flag.Flag) {
	switch f.Name {
	case "confdir":
		config.ConfDir = confdir
	}
}
