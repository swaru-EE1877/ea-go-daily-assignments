package account

import "fmt"

type InsufficientBalanceError struct {
	message             string
	currentBalance      float64
	balanceRequiredToWD float64
}

func (i InsufficientBalanceError) Error() string {
	return fmt.Sprintf("%s. Balance in the account is: %f and required more %f balance to withdraw", i.message, i.currentBalance, i.balanceRequiredToWD)
}
