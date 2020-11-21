package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"vinet/address"
	"vinet/auth"
	"vinet/balance"
	"vinet/board"
	"vinet/coinin"
	"vinet/deposit"
	"vinet/market"
	"vinet/menu"
	"vinet/message"
	"vinet/sendchildorder"
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
			productCode, err := market.SelectProductCode()
			if err != nil {
				log.Println(err)
				break
			}

			fmt.Println(message.GetOrderType())
			fmt.Print(message.GetInputLine())
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			fmt.Println("")
			childOrderType := scanner.Text()
			if childOrderType != sendchildorder.TypeLimit && childOrderType != sendchildorder.TypeMarket {
				fmt.Println("wrong")
				break
			}

			fmt.Println(message.GetSide())
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			side := scanner.Text()
			if side != sendchildorder.SideBuy && side != sendchildorder.SideSell {
				fmt.Println("wrong")
				break
			}

			var price float64
			if childOrderType == sendchildorder.TypeLimit {
				fmt.Println(message.GetPrice())
				fmt.Print(message.GetInputLine())
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				fmt.Println("")
				price, err = strconv.ParseFloat(scanner.Text(), 64)
				if err != nil || price < 0.0 {
					fmt.Println("wrong")
					break
				}
			}

			fmt.Println(message.GetSize())
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			size, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("wrong")
				break
			}

			fmt.Println(message.GetMinuteToExpire())
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			miniteToExpire, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("wrong")
				break
			}

			fmt.Println(message.GetTimeInForce())
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			timeInForce := scanner.Text()
			if timeInForce != sendchildorder.TimeInForceGTC &&
				timeInForce != sendchildorder.TimeInForceIOC &&
				timeInForce != sendchildorder.TimeInForceFOK {
				fmt.Println("wrong")
				break
			}

			b := sendchildorder.Body{
				ProductCode:    productCode,
				ChildOrderType: childOrderType,
				Side:           side,
				Price:          price,
				Size:           size,
				MinuteToExpire: miniteToExpire,
				TimeInForce:    timeInForce,
			}

			if err := sendchildorder.Send(b); err != nil {
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
