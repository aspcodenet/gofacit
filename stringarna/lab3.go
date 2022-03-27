package stringarna

import (
	"fmt"
	"strings"
)

func Run3() {
	string1 := "kurt andersson"

	string1 = strings.ToTitle(string1)

	fmt.Println(string1)

	//Variant 2
	string1 = "kurt andersson"
	string2 := ""
	nextShouldBeLarge := true

	for _, ch := range string1 {
		if nextShouldBeLarge {
			string2 += strings.ToUpper(string(ch))
			nextShouldBeLarge = false
		} else {
			string2 += string(ch)
			if ch == ' ' {
				nextShouldBeLarge = true
			}
		}

	}
	fmt.Println(string2)

}
