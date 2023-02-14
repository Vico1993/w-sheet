package main

import (
	"os"
	"testing"
)

func mockUserHomeDir(dir string, err error) {
	userHomeDir = func() (string, error) {
		return dir, err
	}
}

func TestInitConfig(t *testing.T) {
	defer func() {
		os.Remove("./tmp")
	}()

	mockUserHomeDir("./tmp", nil)

	initConfig()

	if _, err := os.Stat("./tmp/.w/data.json"); os.IsNotExist(err) {
		t.Errorf("config file does not exist")
	}
}
