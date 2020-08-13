package main

import (
	"./accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("klaus")
	// Go 에서 struct 를 출력하면 String 메서드를 호출한다.
	fmt.Println(account)

	account.Deposit(1000)
	fmt.Println(account)

	withDrawAccount(&account, 2000)
	withDrawAccount(&account, 500)
}

func withDrawAccount(account **accounts.Account, amount int) {
	err := (*account).Withdraw(amount)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*account)
	}
}
