/*
Example with configuration file.
*/
package main

import (
	"encoding/json"
	"fmt"

	"github.com/crgimenes/goConfig"
	_ "github.com/crgimenes/goConfig/json"
)

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type systemUser struct {
	Name     string `cfg:"name"`
	Password string `cfg:"passwd"`
}

type configTest struct {
	Domain  string
	User    systemUser `cfg:"user"`
	MongoDB mongoDB
}

func main() {
	config := configTest{}

	goConfig.File = "config.json"
	goConfig.PrefixEnv = "EXAMPLE"
	err := goConfig.Parse(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")
	fmt.Println(string(j))
}
