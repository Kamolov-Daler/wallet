package main

import (
	"fmt"

	"github.com/Kamolov-Daler/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	svc.ImportFromFile("data/account.txt")
	fmt.Println(svc.FindAccountByID(2))
}
