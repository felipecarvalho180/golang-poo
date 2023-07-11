package accounts

import (
	"bank/users"
	"fmt"
)

type CheckingAccount struct {
	User                        users.User
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}

func (c *CheckingAccount) Withdraw(value float64) {
	if value <= c.balance && value > 0 {
		c.balance -= value
	} else {
		fmt.Println("Error, insufficient balance or invalid value")
	}
}

func (c *CheckingAccount) Deposit(value float64) {
	if value > 0 {
		c.balance += value
		fmt.Println("Deposit was made with successfully")
	} else {
		fmt.Println("Deposit error, invalid value")
	}
}

func (c *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount) bool {
	if value < c.balance && value > 0 {
		c.balance -= value
		destinationAccount.balance += value
		return true
	} else {
		return false
	}
}
