package main

import (
	"fmt"

	a "github.com/afaferz/bank/accounts"
	o "github.com/afaferz/bank/owners"
)

func main() {
	FerrazAccount := a.CheckAccount{
		Owner: o.Owner{
			Name:       "Ferraz",
			Cpf:        "067.131.331-26",
			Occupation: "Front-end developer",
		},
		Agency: 589,
		Number: 123456,
	}

	FerrazAccountSaving := a.SavingAccount{
		Owner: o.Owner{
			Name:       "Ferraz",
			Cpf:        "067.131.331-26",
			Occupation: "Front-end developer",
		},
		Agency: 590,
		Number: 123457,
	}

	FerrazAccount.Deposit(100)
	FerrazAccountSaving.Deposit(150)
	FerrazAccountSaving.Withdraw(100)
	PayBankSlip(&FerrazAccountSaving, 20)
	fmt.Println(FerrazAccount.GetWithdraw(), FerrazAccountSaving.GetWithdraw())
}

type VerifyAccount interface {
	Withdraw(value float64) bool
}

func PayBankSlip(account VerifyAccount, bankSlipValue float64) {
	account.Withdraw(bankSlipValue)
}
