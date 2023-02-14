package main

import (
	"log"
	"os"
)

// For now the data will come from a json in the root of the project.
// Could be move later in small SQLite DB?
// Or even a JSON but setup by the CLI and not the user.
// Temp for now and see how it goes.

var userHomeDir = os.UserHomeDir

func initConfig() {
	homedir, err := userHomeDir()
	if err != nil {
		homedir = "./"
	}

	path := homedir + "/.w"

	// Check if .w folder exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("Can't create Folder at " + path)
		}
	}

	configFilePath := path + "/data.json"

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
		if err != nil {
			log.Fatal("Can't create config file at " + configFilePath)
		}
		defer file.Close()

		// initialisation of the JSON string
		_, err = file.WriteString("{}")
		if err != nil {
			log.Fatal("Can't initiate JSON config file " + err.Error())
		}
	}

	v.SetConfigFile(configFilePath)

	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
}
