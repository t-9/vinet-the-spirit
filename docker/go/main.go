package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"vinet/board"
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
		case "2":
			fmt.Println(message.GetWhichBoard())
			choices, err := market.PrintChoices()
			if err != nil {
				log.Fatal(err)
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
