package start

import (
	"fmt"
	"math"
)

func Run6() {
	var tal1, tal2 float64

	fmt.Printf("Skriv in tal 1:")
	fmt.Scanln(&tal1)

	fmt.Printf("Skriv in tal 2:")
	fmt.Scanln(&tal2)

	resultat := math.Pow(float64(tal1), float64(tal2))
	fmt.Printf("%f pow %f = %f\n", tal1, tal2, resultat)

	resultat = tal1 * tal2
	fmt.Printf("%f * %f = %f\n", tal1, tal2, resultat)

	resultat = tal1 / tal2
	fmt.Printf("%f / %f = %f\n", tal1, tal2, resultat)

	resultat = tal1 - tal2
	fmt.Printf("%f - %f = %f\n", tal1, tal2, resultat)

	resultat = tal1 + tal2
	fmt.Printf("%f + %f = %f\n", tal1, tal2, resultat)

}
