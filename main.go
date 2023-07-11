package main

import (
	"bank/accounts"
	"bank/users"
	"fmt"
)

func main() {
	// Corrente
	user1 := users.User{
		Name:          "Felipe",
		Documentation: "174.918.745-09",
		Profession:    "Software Developer",
	}

	user2 := users.User{
		Name:          "Thainá",
		Documentation: "01.829.534/0001-98",
		Profession:    "Designer",
	}

	felipeAccount := accounts.CheckingAccount{
		User:          user1,
		AgencyNumber:  123,
		AccountNumber: 123456,
	}

	thainaAccount := accounts.CheckingAccount{
		User:          user2,
		AgencyNumber:  321,
		AccountNumber: 654321,
	}

	fmt.Println("\nConta Corrente")
	fmt.Println("")

	felipeAccount.Deposit(300)
	felipeAccount.Withdraw(100)

	fmt.Println("Felipe balance:", felipeAccount.GetBalance())
	fmt.Println("Thainá balance:", thainaAccount.GetBalance())

	felipeAccount.Transfer(50, &thainaAccount)

	fmt.Println("\nFelipe balance after transfer:", felipeAccount.GetBalance())
	fmt.Println("Thainá balance after transfer:", thainaAccount.GetBalance())

	PayBill(&felipeAccount, 50)
	fmt.Println("Felipe balance after pay the bill:", felipeAccount.GetBalance())

	// Poupança
	user3 := users.User{
		Name:          "Eliane",
		Documentation: "174.918.745-09",
		Profession:    "None",
	}

	elianeAccount := accounts.SavingsAccount{
		User:          user3,
		AgencyNumber:  123,
		AccountNumber: 123456,
	}

	fmt.Println("\nConta Poupança")
	fmt.Println("")

	elianeAccount.Deposit(300)
	elianeAccount.Withdraw(100)

	fmt.Println("Eliane balance:", elianeAccount.GetBalance())

	PayBill(&elianeAccount, 50)

	fmt.Println("Eliane balance after pay the bill:", elianeAccount.GetBalance())
}

func PayBill(account Account, value float64) {
	account.Withdraw(value)
}

type Account interface {
	Withdraw(value float64)
}
