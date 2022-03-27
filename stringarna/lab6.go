package stringarna

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run6() {
	var text string

	fmt.Print("Ange text:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToLower(text)
	text2 := ""

	for i := len(text) - 1; i >= 0; i-- {
		text2 += string(text[i])
	}

	if text2 == text {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Ej palindrom")
	}

}
