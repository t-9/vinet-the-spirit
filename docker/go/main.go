package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
				log.Fatal(err)
			}
		}

		if shouldExit {
			break
		}
		fmt.Println("")
	}

	fmt.Println(message.GetBye())
}
