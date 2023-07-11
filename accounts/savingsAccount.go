package accounts

import (
	"bank/users"
	"fmt"
)

type SavingsAccount struct {
	User                                   users.User
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}

func (c *SavingsAccount) Withdraw(value float64) {
	if value <= c.balance && value > 0 {
		c.balance -= value
	} else {
		fmt.Println("Error, insufficient balance or invalid value")
	}
}

func (c *SavingsAccount) Deposit(value float64) {
	if value > 0 {
		c.balance += value
		fmt.Println("Deposit was made with successfully")
	} else {
		fmt.Println("Deposit error, invalid value")
	}
}
