package ifs

import (
	"fmt"
	"strings"
)

func Run8() {
	var belopp uint
	var rabattInput string
	var harRabatt bool

	fmt.Print("Hur mycket pengar har du?")
	fmt.Scanln(&belopp)

	fmt.Print("Har du rabatt j/n?")
	fmt.Scanln(&rabattInput)
	if strings.ToLower(rabattInput) == "j" {
		harRabatt = true
	}

	if belopp >= 15 && belopp <= 25 {
		if harRabatt {
			fmt.Println("Du kan köpa en liten hamburgare och en pommes frites")
		} else {
			fmt.Println("Du kan köpa en liten hamburgare")
		}
	} else if belopp > 25 && belopp <= 50 {
		if harRabatt {
			fmt.Println("Du kan köpa en storn hamburgare och en pommes frites")
		} else {
			fmt.Println("Du kan köpa en stor hamburgare")
		}
	} else if belopp > 60 || (belopp >= 50 && belopp <= 60 && harRabatt) {
		fmt.Println("Du kan köpa ett meal med dryck")
	}

}
