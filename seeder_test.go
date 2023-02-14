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
	viper.SetConfigFile("./test.json")
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
	err := ioutil.WriteFile("./seeder.json", seederJSON, 0644)
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
		from: asset{
			quantiy:  800,
			code:     "USD",
			isCrypto: true,
		},
		to: asset{
			quantiy:  1000,
			code:     "CAD",
			isCrypto: false,
		},
		date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	if transactions[0].from != expected.from || transactions[0].to != expected.to || transactions[0].date != expected.date {
		t.Errorf("Expected %v, but got %v", expected, transactions[0])
		return
	}

	// Clean up
	err = os.Remove("./seeder.json")
	if err != nil {
		t.Errorf("Failed to remove test seeder file: %s", err)
	}
}

func TestSeederFileNotFound(t *testing.T) {
	// Load transactions from seeder
	transactions, err := loadSeeder()

	assert.Nil(t, transactions, "Expected nil instead of transactions")
	assert.EqualError(t, err, "open ./seeder.json: no such file or directory", "An file not found error should be triggered")
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
	err := ioutil.WriteFile("./seeder.json", seederJSON, 0644)
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
	err := ioutil.WriteFile("./seeder.json", seederJSON, 0644)
	if err != nil {
		t.Errorf("Failed to create test seeder file: %s", err)
		return
	}

	// Load transactions from seeder
	transactions, err := loadSeeder()

	assert.Nil(t, transactions, "Expected nil instead of transactions")
	assert.EqualError(t, err, "No transactions found", "An transactions not found error should be triggered")
}
