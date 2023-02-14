package main

import (
	"time"

	"github.com/google/uuid"
)

const CURRENCY_CODE = "CAD"

type fiat struct {
	Value    float64
	Currency string
}

type asset struct {
	Quantiy  float64
	Code     string
	IsCrypto bool
}

type transaction struct {
	Id        string
	From      asset
	To        asset
	Reference fiat
	Date      time.Time
}

// Create a new Transaction from exchange information
func newTransaction(from asset, to asset, date time.Time) *transaction {
	// Default reference
	ref := fiat{
		Currency: CURRENCY_CODE,
		Value:    0,
	}

	// Purchase
	if !from.IsCrypto && from.Code == CURRENCY_CODE {
		ref = fiat{
			Currency: CURRENCY_CODE,
			Value:    from.Quantiy,
		}
	}

	// Withdrawal
	if !to.IsCrypto && to.Code == CURRENCY_CODE {
		ref = fiat{
			Currency: CURRENCY_CODE,
			Value:    to.Quantiy,
		}
	}

	return &transaction{
		Id:        uuid.New().String(),
		From:      from,
		To:        to,
		Date:      date,
		Reference: ref,
	}
}
