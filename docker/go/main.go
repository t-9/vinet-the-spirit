package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"vinet/auth"
	"vinet/balance"
	"vinet/board"
	"vinet/coinin"
	"vinet/deposit"
	"vinet/market"
	"vinet/message"
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
			fmt.Println(message.GetWhichBoard())
			choices, err := market.PrintChoices()
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Print(message.GetInputLine())
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			fmt.Println("")

			c, cerr := strconv.Atoi(scanner.Text())
			if cerr != nil || c >= len(choices) || c < 0 {
				fmt.Println(message.GetWrongChoice())
				break
			}

			if err := board.PrintList(choices[c]); err != nil {
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
			scanner = bufio.NewScanner(os.Stdin)
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
		}

		if shouldExit {
			break
		}
		fmt.Println("")
	}

	fmt.Println(message.GetBye())
}
