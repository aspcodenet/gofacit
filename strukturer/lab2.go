package strukturer

import (
	"fmt"
	"strings"
)

func Run2() {
	allaNamn := make([]string, 1, 10)

	for {
		var namn string
		fmt.Print("Ange ett namn:")
		fmt.Scanln(&namn)

		allaNamn = append(allaNamn, namn)

		var cont string
		fmt.Print("Vill du fortsätta y?")
		fmt.Scanln(&cont)
		cont = strings.ToLower(cont)
		if cont == "y" || cont == "yes" {
			continue
		}
		break

	}

	longest := allaNamn[0]
	for _, val := range allaNamn {
		if len(val) > len(longest) {
			longest = val
		}
	}

	fmt.Printf("Longest is %s\n", longest)

}
