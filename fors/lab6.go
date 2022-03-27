package fors

import (
	"fmt"
	"math/rand"
	"strings"
)

func Run6() {

	const min = 1
	const max = 6
	for {

		tal1 := rand.Intn(max-min+1) + min
		tal2 := rand.Intn(max-min+1) + min

		fmt.Printf("%d och %d\n", tal1, tal2)

		var cont string
		fmt.Print("Vill du fortsätta y?")
		fmt.Scanln(&cont)
		cont = strings.ToLower(cont)
		if cont == "y" || cont == "yes" {
			continue
		}
		break
	}
}
