package main

import (
	"fmt"
	"time"
)

func main() {
	tr := newTransaction(asset{
		quantiy:  1,
		code:     "BTC",
		isCrypto: true,
	}, asset{
		quantiy:  100,
		code:     "CAD",
		isCrypto: false,
	}, time.Now().AddDate(1993, 10, 03))

	fmt.Println("Hello:", tr.id)
}
