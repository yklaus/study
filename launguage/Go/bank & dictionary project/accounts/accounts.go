package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw money.")

// NewAccount creates Account
// constructor in swift / kotlin
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// extension function like swift / kotlin
// Go 에서는 Receiver Function 이라고 부른다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount on your account
// Go 에는 try-catch 가 없다
// 그래서 error 처리를 직접해줘야 한다.
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// String : fmt.Println(struct) 를 하면 호출되는 기본 메서드
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
