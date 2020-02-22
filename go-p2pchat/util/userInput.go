package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
	}
}
