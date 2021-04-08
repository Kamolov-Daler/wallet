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
