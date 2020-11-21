package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"vinet/message"
	"vinet/util"
)

func SelectMenu() (int64, error) {
	menuItemList := []int64{
		ShowMarkets,
		ShowBoard,
		RegisterAccessKey,
		ShowBalance,
		ShowCoinIn,
		ShowDeposit,
		ShowAddress,
		SendChildOrder,
		Exit,
	}

	fmt.Println(message.GetWhatDoYouDo())
	for _, i := range menuItemList {
		fmt.Printf("%d. %s\n", i, getMenuItemMessage(i))
	}
	fmt.Print(message.GetInputLine())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("")

	c, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil || !util.ContainsInt64(menuItemList, c) {
		return -1, fmt.Errorf(message.GetWrongChoice())
	}
	return c, err
}

func getMenuItemMessage(n int64) string {
	switch n {
	case ShowMarkets:
		return message.GetShowMarkets()
	case ShowBoard:
		return message.GetShowBoard()
	case RegisterAccessKey:
		return message.GetRegisterAccessKey()
	case ShowBalance:
		return message.GetShowBalance()
	case ShowCoinIn:
		return message.GetShowCoinIn()
	case ShowDeposit:
		return message.GetShowDeposit()
	case ShowAddress:
		return message.GetShowAddress()
	case SendChildOrder:
		return message.GetSendChildOrder()
	case Exit:
		return message.GetExit()
	}
	return ""
}
