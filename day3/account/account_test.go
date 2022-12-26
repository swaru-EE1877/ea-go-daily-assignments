package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}
func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}
func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}
func TestThrowErrorWhenInsufficientBalance(t *testing.T) {
	acc := Account{balance: 500}

	err := acc.Withdraw(700)

	insufficientBalanceError := err.(InsufficientBalanceError)

	assert.Equal(t, insufficientBalanceError, InsufficientBalanceError{message: "Insufficient balance", currentBalance: 500, balanceRequiredToWD: -200})
}
