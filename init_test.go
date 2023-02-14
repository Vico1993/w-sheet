package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockUserHomeDir(dir string, err error) {
	userHomeDir = func() (string, error) {
		return dir, err
	}
}

func TestInitConfig(t *testing.T) {
	defer func() {
		os.RemoveAll("./tmp")
	}()

	mockUserHomeDir("./tmp", nil)

	err := initConfig()
	assert.Nil(t, err, "Error should be nul")

	if _, err = os.Stat("./tmp/.w/data.json"); os.IsNotExist(err) {
		t.Errorf("config file does not exist")
	}
}

func TestDefaultHomeDir(t *testing.T) {
	defer func() {
		os.RemoveAll(".w")
	}()

	mockUserHomeDir("./tmp", errors.New("Ooops"))

	err := initConfig()
	assert.Nil(t, err, "Error should be nul")

	if _, err := os.Stat("./.w/data.json"); os.IsNotExist(err) {
		t.Errorf("config file does not exist")
	}
}
