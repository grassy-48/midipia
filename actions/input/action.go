package input

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"
	cli "gopkg.in/urfave/cli.v1"
)

// Input push keyboard
func Input(c *cli.Context) error {
	fmt.Println("a:C ~ k:C & space: crotchet rest & Enter:submit")
	if terminal.IsTerminal(0) {
		var str string
		// waiting input
		fmt.Scan(&str)
		fmt.Println("渡された標準入力を出力:", str)
	} else {
		// get pipe score
		// ex: echo "aagghhg ffddssa" | go run main.go input-score
		str, _ := ioutil.ReadAll(os.Stdin)
		fmt.Printf("input: %s", string(str))
	}

	return nil
}
