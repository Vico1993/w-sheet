package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTransaction(t *testing.T) {
	transac := newTransaction(asset{
		Quantiy:  1,
		Code:     "BTC",
		IsCrypto: true,
	}, asset{
		Quantiy:  100,
		Code:     "CAD",
		IsCrypto: false,
	}, time.Date(1993, 10, 3, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, float64(100), transac.Reference.Value)
}

func TestDefaultTransactionRef(t *testing.T) {
	transac := newTransaction(asset{
		Quantiy:  1,
		Code:     "BTC",
		IsCrypto: true,
	}, asset{
		Quantiy:  100,
		Code:     "ETH",
		IsCrypto: true,
	}, time.Date(1993, 10, 3, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, float64(0), transac.Reference.Value)
}

func TestRefWithFromCurrency(t *testing.T) {
	transac := newTransaction(asset{
		Quantiy:  100,
		Code:     "CAD",
		IsCrypto: false,
	}, asset{
		Quantiy:  1,
		Code:     "BTC",
		IsCrypto: true,
	}, time.Date(1993, 10, 3, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, float64(100), transac.Reference.Value)
}
