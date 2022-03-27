package stringarna

import (
	"fmt"
	"strings"
)

func Run5() {
	var email string

	fmt.Print("Ange epostaddress:")
	fmt.Scanln(&email)

	atPos := strings.Index(email, "@")
	lastPointPos := strings.LastIndex(email, ".")
	charsAfter := len(email) - lastPointPos - 1

	if atPos > 1 && (charsAfter == 2 || charsAfter == 3) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not valid")
	}

}
