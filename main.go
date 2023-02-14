package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var v = viper.GetViper()

func main() {
	// Init configuration file
	initConfig()

	// Load seeder if there is any
	loadSeeder()

	var transactions []transaction

	err := v.UnmarshalKey("transactions", &transactions)
	if err != nil {
		log.Fatalln("Error loading operations: ", err.Error())
	}

	for _, transaction := range transactions {
		fmt.Println("Transaction", transaction.date, " - ", transaction.id)
	}
}
