package sendchildorder

import "fmt"

func (r Response) String() string {
	return r.ChildOrderAcceptanceID
}

func Send(b Body) error {
	s, err := sendChildOrder(b)
	if err != nil {
		return err
	}

	fmt.Println("ChildOrderAcceptanceID")
	fmt.Println(s)
	return nil
}
