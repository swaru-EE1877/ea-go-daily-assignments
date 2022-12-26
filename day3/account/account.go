package account

// TDD Bank Account app

type Account struct {
	balance float64
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	if acc.balance < amount {
		return InsufficientBalanceError{"Insufficient balance", acc.balance, acc.balance - amount}
	}
	acc.balance -= amount
	return nil
}
