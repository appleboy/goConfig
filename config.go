package goConfig

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/crgimenes/goConfig/goEnv"
	"github.com/crgimenes/goConfig/goFlags"
)

// Settings default
type Settings struct {
	// Path sets default config path
	Path string
	// File name of default config file
	File string
	// FileRequired config file required
	FileRequired bool
}

const tag = "cfg"
const tagDefault = "cfgDefault"

// Setup Pointer to internal variables
var Setup *Settings

func init() {
	Setup = &Settings{
		Path:         "./",
		File:         "config.json",
		FileRequired: false,
	}
}

// Parse configuration
func Parse(config interface{}) (err error) {

	err = LoadJSON(config)
	if err != nil {
		return
	}

	goEnv.Setup(tag, tagDefault)
	err = goEnv.Parse(config)
	if err != nil {
		return
	}

	goFlags.Setup(tag, tagDefault)
	goFlags.Preserve = true
	err = goFlags.Parse(config)
	if err != nil {
		return
	}

	return
}

// LoadJSON config file
func LoadJSON(config interface{}) (err error) {
	configFile := Setup.Path + Setup.File
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !Setup.FileRequired {
		err = nil
		return
	} else if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	return
}

// Save config file
func Save(config interface{}) (err error) {
	_, err = os.Stat(Setup.Path)
	if os.IsNotExist(err) {
		os.Mkdir(Setup.Path, 0700)
	} else if err != nil {
		return
	}

	configFile := Setup.Path + Setup.File

	_, err = os.Stat(configFile)
	if err != nil {
		return
	}

	b, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return
	}

	err = ioutil.WriteFile(configFile, b, 0644)
	if err != nil {
		return
	}
	return
}
