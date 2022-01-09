package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func Apply(token *string, prefix *string) {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	jsonConfig, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var config Config

	json.Unmarshal(jsonConfig, &config)

	*token = config.Token
	*prefix = config.Prefix
}
