package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type seederJsonReference struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type seederJson struct {
	Quantity     string              `json:"quantity"`
	Code         string              `json:"code"`
	FromCode     string              `json:"from_code"`
	FromQuantity string              `json:"from_quantity"`
	Date         string              `json:"date"`
	Reference    seederJsonReference `json:"reference"`
}

var seederFileName = "./seeder.json"

func loadSeeder() ([]transaction, error) {
	file, err := ioutil.ReadFile(seederFileName)
	if err != nil {
		return nil, err
	}

	var data []seederJson
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	// Parse data and transform into transactions
	var transactions []transaction
	for _, operation := range data {
		quantity, _ := strconv.ParseFloat(operation.Quantity, 64)
		quantityFrom, _ := strconv.ParseFloat(operation.FromQuantity, 64)
		date, _ := time.Parse("2006-01-02", operation.Date)

		transactions = append(transactions, *newTransaction(asset{
			Quantiy:  quantityFrom,
			Code:     operation.FromCode,
			IsCrypto: operation.FromCode != "CAD",
		}, asset{
			Quantiy:  quantity,
			Code:     operation.Code,
			IsCrypto: operation.Code != "CAD"}, date))
	}

	if len(transactions) == 0 {
		return nil, errors.New("No transactions found")
	}

	// For now just erase what we have.@
	v.Set("transactions", transactions)

	err = v.WriteConfig()
	if err != nil {
		return nil, err
	}

	// clean seeder file
	err = os.Remove(seederFileName)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
