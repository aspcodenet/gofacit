package fors

import (
	"fmt"
	"strings"
)

func Run3() {
	var tal1, tal2 int

	for {
		fmt.Print("Ange tal1:")
		fmt.Scanln(&tal1)
		fmt.Print("Ange tal2:")
		fmt.Scanln(&tal2)

		fmt.Printf("%d+%d=%d\n", tal1, tal2, tal1+tal2)

		var cont string
		fmt.Print("Vill du fortsätta (j/n)?")
		fmt.Scanln(&cont)
		if strings.ToLower(cont) != "j" {
			break
		}
	}

}
