package wallet

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Kamolov-Daler/wallet/pkg/wallet"
)

func TestService_FindAccountById_find(t *testing.T) {
	svc := &wallet.Service{}
	_, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := svc.FindAccountByID(3)

	expected, err := "account not found", nil

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}

}
