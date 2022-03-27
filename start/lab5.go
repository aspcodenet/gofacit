package start

import (
	"fmt"
	"math"
)

func Run5() {
	var tal1, tal2 int32

	fmt.Printf("Skriv in tal 1:")
	fmt.Scanln(&tal1)

	fmt.Printf("Skriv in tal 2:")
	fmt.Scanln(&tal2)

	resultat := int32(math.Pow(float64(tal1), float64(tal2)))
	fmt.Printf("%d pow %d = %d\n", tal1, tal2, resultat)

	resultat = tal1 * tal2
	fmt.Printf("%d * %d = %d\n", tal1, tal2, resultat)

	resultat = tal1 / tal2
	fmt.Printf("%d / %d = %d\n", tal1, tal2, resultat)

	resultat = tal1 - tal2
	fmt.Printf("%d - %d = %d\n", tal1, tal2, resultat)

	resultat = tal1 + tal2
	fmt.Printf("%d + %d = %d\n", tal1, tal2, resultat)

}
