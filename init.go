package main

import (
	"os"
)

// For now the data will come from a json in the root of the project.
// Could be move later in small SQLite DB?
// Or even a JSON but setup by the CLI and not the user.
// Temp for now and see how it goes.

var userHomeDir = os.UserHomeDir

func initConfig() error {
	homedir, err := userHomeDir()
	if err != nil {
		homedir = "./"
	}

	path := homedir + "/.w"

	// Check if .w folder exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	configFilePath := path + "/data.json"

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// initialisation of the JSON string
		_, err = file.WriteString("{}")
		if err != nil {
			return err
		}
	}

	v.SetConfigFile(configFilePath)

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
