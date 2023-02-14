package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func before() {
	seederFileName = "./test-seeder.json"

	viper.SetConfigFile("./test.json")

	defer func() {

		os.Remove("./test.json")
		os.Remove("./test-seeder.json")
	}()
}

func TestLoadSeeder(t *testing.T) {
	before()

	// Create a test JSON file for seeder
	seederJSON := []byte(`[
		{
			"code": "CAD",
			"quantity": "1000",
			"from_code": "USD",
			"from_quantity": "800",
			"date": "2022-01-01"
		}
	]`)
	err := ioutil.WriteFile(seederFileName, seederJSON, 0644)
	if err != nil {
		t.Errorf("Failed to create test seeder file: %s", err)
		return
	}

	// Load transactions from seeder
	transactions, err := loadSeeder()
	assert.Nil(t, err, "Error should be nil")

	// Check the number of transactions
	if len(transactions) != 1 {
		t.Errorf("Expected 1 transaction, but got %d", len(transactions))
		return
	}

	// Check the contents of the transaction
	expected := transaction{
		From: asset{
			Quantiy:  800,
			Code:     "USD",
			IsCrypto: true,
		},
		To: asset{
			Quantiy:  1000,
			Code:     "CAD",
			IsCrypto: false,
		},
		Date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	if transactions[0].From != expected.From || transactions[0].To != expected.To || transactions[0].Date != expected.Date {
		t.Errorf("Expected %v, but got %v", expected, transactions[0])
		return
	}
}

func TestSeederFileNotFound(t *testing.T) {
	// Load transactions from seeder
	transactions, err := loadSeeder()

	assert.Nil(t, transactions, "Expected nil instead of transactions")
	assert.EqualError(t, err, "open "+seederFileName+": no such file or directory", "An file not found error should be triggered")
}

func TestSeederIncorrect(t *testing.T) {
	before()

	// Create a test JSON file for seeder
	seederJSON := []byte(`[
		{
			"code": "CAD"
			"quantity": "1000"
			"from_code": "USD",
			"from_quantity": "800",
			"date": "2022-01-01"
		}
	]`)
	err := ioutil.WriteFile(seederFileName, seederJSON, 0644)
	if err != nil {
		t.Errorf("Failed to create test seeder file: %s", err)
		return
	}

	// Load transactions from seeder
	transactions, err := loadSeeder()

	assert.Nil(t, transactions, "Expected nil instead of transactions")
	assert.EqualError(t, err, `invalid character '"' after object key:value pair`, "An invalid character error should be triggered")
}

func TestEmptySeeder(t *testing.T) {
	before()

	// Create a test JSON file for seeder
	seederJSON := []byte(`[]`)
	err := ioutil.WriteFile(seederFileName, seederJSON, 0644)
	if err != nil {
		t.Errorf("Failed to create test seeder file: %s", err)
		return
	}

	// Load transactions from seeder
	transactions, err := loadSeeder()

	assert.Nil(t, transactions, "Expected nil instead of transactions")
	assert.EqualError(t, err, "No transactions found", "An transactions not found error should be triggered")
}
