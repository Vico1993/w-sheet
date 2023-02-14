package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func loadSeeder() []transaction {
	file, err := ioutil.ReadFile("./seeder.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var data []seederJson
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Parse data and transform into transactions
	var transactions []transaction
	for _, operation := range data {
		quantity, _ := strconv.ParseFloat(operation.Quantity, 64)
		quantityFrom, _ := strconv.ParseFloat(operation.FromQuantity, 64)
		date, _ := time.Parse("2006-01-02", operation.Date)

		transactions = append(transactions, *newTransaction(asset{
			quantiy:  quantityFrom,
			code:     operation.FromCode,
			isCrypto: operation.FromCode != "CAD",
		}, asset{
			quantiy:  quantity,
			code:     operation.Code,
			isCrypto: operation.Code != "CAD"}, date))
	}

	if len(transactions) == 0 {
		fmt.Println("Couldn't load transactions")
	}

	// For now just erase what we have.@
	v.Set("transactions", transactions)

	err = v.WriteConfig()
	if err != nil {
		log.Fatalln("Error saving transactions: ", err.Error())
	}

	return transactions
}
