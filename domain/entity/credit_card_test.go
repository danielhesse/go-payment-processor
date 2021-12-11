package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	// Invalid CreditCard Number
	_, err := NewCreditCard("40000000000000000", "Daniel Hessel", 12, 2024, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	// Valid CreditCard Number
	_, err = NewCreditCard("5170630238686544", "Daniel Hessel", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5170630238686544", "Daniel Hessel", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5170630238686544", "Daniel Hessel", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5170630238686544", "Daniel Hessel", 11, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("5170630238686544", "Daniel Hessel", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
