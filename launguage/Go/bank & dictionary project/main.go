package main

import (
	"./dictionary"
	"fmt"
)

func main() {
	//account := accounts.NewAccount("klaus")
	//// Go 에서 struct 를 출력하면 String 메서드를 호출한다.
	//fmt.Println(account)
	//
	//account.Deposit(1000)
	//fmt.Println(account)
	//
	//withDrawAccount(&account, 2000)
	//withDrawAccount(&account, 500)

	dict := dictionary.Dictionary{"first": "First word"}
	//definition, err := dict.Search("first")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(definition)
	//}

	word := "second"
	def := "Second word"
	err := dict.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}

	second, _ := dict.Search(word)
	fmt.Println("found", word, "definition: ", second)

	//err2 := dict.Add(word, def)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}

	err3 := dict.Update("second", "Change Word")
	if err3 != nil {
		fmt.Println(err3)
	}

	//second, _ = dict.Search(word)
	//fmt.Println("found", word, "definition: ", second)
	fmt.Println(dict)

	dict.Delete("second")
	fmt.Println(dict)
}

//func withDrawAccount(account **accounts.Account, amount int) {
//	err := (*account).Withdraw(amount)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(*account)
//	}
//}
