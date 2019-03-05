package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var envkey map[string]string

// EnvSetup - load env from json
func EnvSetup() {

	jsonFile, err := os.Open("config/env.json")
	defer jsonFile.Close()
	if err != nil {
		panic("failed to open json file: " + err.Error())
	}

	rawJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		panic("failed to read json file : " + err.Error())
	}

	err = json.Unmarshal(rawJSON, &envkey)
	if err != nil {
		panic("failed to unmarhsall env.json : " + err.Error())
	}

	for key, value := range envkey {
		os.Setenv(key, value)
	}
}
