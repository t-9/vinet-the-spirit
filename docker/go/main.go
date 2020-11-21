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
	"vinet/message"
	"vinet/sendchildorder"
)

func main() {
	fmt.Println(message.GetHello())
	fmt.Println("")

	for {
		fmt.Print(message.GetMenu())
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fmt.Println("")

		shouldExit := false
		switch scanner.Text() {
		case "0":
			shouldExit = true
		case "1":
			if err := market.PrintList(); err != nil {
				log.Println(err)
			}
		case "2":
			productCode, err := market.SelectProductCode()
			if err != nil {
				log.Println(err)
				break
			}
			if err := board.PrintList(productCode); err != nil {
				log.Println(err)
			}
		case "3":
			fmt.Println(message.GetAPIKey())
			fmt.Print(message.GetInputLine())
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			fmt.Println("")
			auth.SetAccessKey(scanner.Text())

			fmt.Println(message.GetAPISecret())
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			if err := auth.SetAccessSecret(scanner.Text()); err != nil {
				log.Println(err)
			}
		case "4":
			if err := balance.PrintList(); err != nil {
				log.Println(err)
			}
		case "5":
			if err := coinin.PrintList(); err != nil {
				log.Println(err)
			}
		case "6":
			if err := deposit.PrintList(); err != nil {
				log.Println(err)
			}
		case "7":
			if err := address.PrintList(); err != nil {
				log.Println(err)
			}
		case "8":
			productCode, err := market.SelectProductCode()
			if err != nil {
				log.Println(err)
				break
			}

			fmt.Println("childOrderType?")
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			childOrderType := scanner.Text()
			if childOrderType != "LIMIT" && childOrderType != "MARKET" {
				fmt.Println("wrong")
				break
			}

			fmt.Println("side?")
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			side := scanner.Text()
			if side != "BUY" && side != "SELL" {
				fmt.Println("wrong")
				break
			}

			var price float64
			if childOrderType == "LIMIT" {
				fmt.Println("price?")
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

			fmt.Println("size?")
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			size, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("wrong")
				break
			}

			fmt.Println("minite_to_expire?")
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			miniteToExpire, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("wrong")
				break
			}

			fmt.Println("time_in_force?")
			fmt.Print(message.GetInputLine())
			scanner.Scan()
			fmt.Println("")
			timeInForce := scanner.Text()
			if timeInForce != "GTC" && timeInForce != "IOC" && timeInForce != "FOK" {
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
