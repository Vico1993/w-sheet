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
	_, err := loadSeeder()
	if err != nil {
		fmt.Println("Couldn't load the transactions", err.Error())
		return
	}

	var transactions []transaction
	err = v.UnmarshalKey("transactions", &transactions)
	if err != nil {
		log.Fatalln("Error loading operations: ", err.Error())
	}

	for _, transaction := range transactions {
		fmt.Println("Transaction", transaction.date, " - ", transaction.id)
	}
}
