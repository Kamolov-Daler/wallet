package main

import (
	"fmt"

	"github.com/Kamolov-Daler/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePosititve:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}
	acc, err := svc.FindAccountByID(account.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc)
}
