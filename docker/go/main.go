package main

import (
	"fmt"
	"log"

	"vinet/address"
	"vinet/auth"
	"vinet/balance"
	"vinet/board"
	sendchild "vinet/childorder/send"
	"vinet/coinin"
	"vinet/deposit"
	"vinet/market"
	"vinet/menu"
	"vinet/message"
	cancelparent "vinet/parentorder/cancel"
	getparent "vinet/parentorder/get"
	sendparent "vinet/parentorder/send"
)

func main() {
	fmt.Println(message.GetHello())
	fmt.Println("")

	for {
		c, err := menu.SelectMenu()
		if err != nil {
			fmt.Println(err)
			fmt.Println("")
			continue
		}

		shouldExit := false
		switch c {
		case menu.Exit:
			shouldExit = true
		case menu.ShowMarkets:
			if err := market.PrintList(); err != nil {
				log.Println(err)
			}
		case menu.ShowBoard:
			productCode, err := market.SelectProductCode()
			if err != nil {
				log.Println(err)
				break
			}
			if err := board.PrintList(productCode); err != nil {
				log.Println(err)
			}
		case menu.RegisterAccessKey:
			if err := auth.RegisterAccessKey(); err != nil {
				log.Println(err)
			}
		case menu.ShowBalance:
			if err := balance.PrintList(); err != nil {
				log.Println(err)
			}
		case menu.ShowCoinIn:
			if err := coinin.PrintList(); err != nil {
				log.Println(err)
			}
		case menu.ShowDeposit:
			if err := deposit.PrintList(); err != nil {
				log.Println(err)
			}
		case menu.ShowAddress:
			if err := address.PrintList(); err != nil {
				log.Println(err)
			}
		case menu.SendChildOrder:
			if err := sendchild.Order(); err != nil {
				log.Println(err)
			}
		case menu.SendParentOrder:
			if err := sendparent.Order(); err != nil {
				log.Println(err)
			}
		case menu.ShowParentOrder:
			if err := getparent.PrintList(); err != nil {
				log.Println(err)
			}
		case menu.CancelParentOrder:
			if err := cancelparent.Cancel(); err != nil {
				log.Println(err)
			}
		}

		if shouldExit {
			break
		}
		fmt.Println("")
	}

	fmt.Println(message.GetBye())
}
