package accounts

import (
	"fmt"

	o "github.com/afaferz/bank/owners"
)

type CheckAccount struct {
	Owner          o.Owner
	Agency, Number int64
	balance        float64
}

func (c *CheckAccount) Deposit(value float64) bool {
	canDeposit := value > 0
	if canDeposit {
		c.balance += value
		// fmt.Printf("R$ %.2f has deposit in your account!", value)
		// fmt.Println("-----------")
		// fmt.Printf("New BALANCE: %.2f", c.balance)
		return true
	} else {
		fmt.Println("Enter with a valid value to deposit")
		return false
	}
}

func (c *CheckAccount) Withdraw(value float64) bool {
	canWithdraw := value > 0 && value <= c.balance
	if canWithdraw {
		c.balance -= value
		fmt.Printf("Successful withdraw!")
		// fmt.Println("-----------")
		// fmt.Printf("New BALANCE: %.2f", c.balance)
		return true
	} else {
		fmt.Printf("You dont have balance!")
		// fmt.Println("-----------")
		// fmt.Printf("Your BALANCE: %.2f", c.balance)
		return false
	}
}

func (c *CheckAccount) Transfer(value float64, fromAccount CheckAccount) bool {
	canTransfer := value < c.balance
	if canTransfer {
		c.balance -= value
		fromAccount.Deposit(value)
		fmt.Println("Successful transfer")
		return true
	} else {
		return false
	}
}

// GETTERS

func (c *CheckAccount) GetWithdraw() float64 {
	return c.balance
}
