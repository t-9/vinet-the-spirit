package board

import (
	"fmt"
	"strings"
)

func (o Order) String() string {
	return fmt.Sprintf("%f, %f", o.Price, o.Size)
}

func stringOrderList(list []Order) string {
	s := make([]string, len(list), len(list))
	for i, o := range list {
		s[i] = o.String()
	}
	return strings.Join(s, "\n")
}

func (b Board) String() string {
	return fmt.Sprintf(`mid_price: %f
Bids
Price, Size
%s

Asks
Price, Size
%s
`, b.MidPrice, stringOrderList(b.Bids), stringOrderList(b.Asks))
}

func PrintList(code string) error {
	board, err := getBoard(code)
	if err != nil {
		return err
	}

	fmt.Println(board)
	return nil
}
