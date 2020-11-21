package auth

import (
	"bufio"
	"fmt"
	"os"

	"vinet/message"
)

// RegisterAccessKey registers an access key and an access secret.
func RegisterAccessKey() error {
	fmt.Println(message.GetAPIKey())
	fmt.Print(message.GetInputLine())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("")
	setAccessKey(scanner.Text())

	fmt.Println(message.GetAPISecret())
	fmt.Print(message.GetInputLine())
	scanner.Scan()
	fmt.Println("")
	return setAccessSecret(scanner.Text())
}
