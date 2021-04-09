package wallet

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Kamolov-Daler/wallet/pkg/types"
)

func TestService_FindAccountById_found(t *testing.T) {
	svc := Service{}
	_, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := svc.FindAccountByID(1)
	expected := &types.Account{
		ID:      1,
		Phone:   "+992000000001",
		Balance: 0,
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestService_FindAccountById_NotFound(t *testing.T) {
	svc := Service{}
	_, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := svc.FindAccountByID(3)

	if result != nil {
		t.Error("result nil!")
	}
}

func TestService_FindPaymentById_found(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		switch err {
		case ErrAmountMustBePosititve:
			fmt.Println("Сумма должна быть положительной")
		case ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}

	payment, err := svc.Pay(account.ID, 10, "mobile")

	if payment == nil {
		return
	}
	result, err := svc.FindPaymentByID(payment.ID)
	expected := &types.Payment{
		ID:        payment.ID,
		AccountID: account.ID,
		Amount:    10,
		Category:  "mobile",
		Status:    types.PaymentStatusInProgress,
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestService_FindPaymentByID_NotFound(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		switch err {
		case ErrAmountMustBePosititve:
			fmt.Println("Сумма должна быть положительной")
		case ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}

	payment, err := svc.Pay(account.ID, 10, "mobile")

	if payment == nil {
		return
	}

	result, err := svc.FindPaymentByID("asdasd")

	if result != nil {
		t.Error("result nil!")
	}
}

func TestService_Reject_found(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		switch err {
		case ErrAmountMustBePosititve:
			fmt.Println("Сумма должна быть положительной")
		case ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}

	payment, err := svc.Pay(account.ID, 10, "mobile")

	if payment == nil {
		return
	}

	result := svc.Reject(payment.ID)

	if result != nil {
		t.Error("result nil!")
	}
}

func TestService_Reject_NotFound(t *testing.T) {
	svc := Service{}
	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 1000)
	if err != nil {
		switch err {
		case ErrAmountMustBePosititve:
			fmt.Println("Сумма должна быть положительной")
		case ErrAccountNotFound:
			fmt.Println("Аккаунт пользователя не найден")
		}
		return
	}

	payment, err := svc.Pay(account.ID, 10, "mobile")

	if payment == nil {
		return
	}

	result := svc.Reject("asdasdasd")

	if result == nil {
		t.Error("result nil!")
	}
}
