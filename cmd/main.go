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

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePosititve:
			fmt.Println("Сумма должна быть положительной")
		case wallet.ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}

	payment, err := svc.Pay(account.ID, 10, "mobile")

	if payment == nil {
		return
	}

	pay := svc.Reject(payment.ID)

	fmt.Println(pay)
}
