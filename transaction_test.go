package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTransaction(t *testing.T) {
	transac := newTransaction(asset{
		quantiy:  1,
		code:     "BTC",
		isCrypto: true,
	}, asset{
		quantiy:  100,
		code:     "CAD",
		isCrypto: false,
	}, time.Now().AddDate(1993, 10, 03))

	assert.Equal(t, float64(100), transac.reference.value)
}
