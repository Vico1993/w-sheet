package main

import (
	"time"

	"github.com/google/uuid"
)

const CURRENCY_CODE = "CAD"

type fiat struct {
	value    int64
	currency string
}

type asset struct {
	quantiy  int64
	code     string
	isCrypto bool
}

type transaction struct {
	id        string
	from      asset
	to        asset
	reference fiat
	date      time.Time
}

// Create a new Transaction from exchange information
func newTransaction(from asset, to asset, date time.Time) *transaction {
	// Default reference
	ref := fiat{
		currency: CURRENCY_CODE,
		value:    0,
	}

	// Purchase
	if !from.isCrypto && from.code == CURRENCY_CODE {
		ref = fiat{
			currency: CURRENCY_CODE,
			value:    from.quantiy,
		}
	}

	// Withdrawal
	if !to.isCrypto && to.code == CURRENCY_CODE {
		ref = fiat{
			currency: CURRENCY_CODE,
			value:    to.quantiy,
		}
	}

	return &transaction{
		id:        uuid.New().String(),
		from:      from,
		to:        to,
		date:      date,
		reference: ref,
	}
}
