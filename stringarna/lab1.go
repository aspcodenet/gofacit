package stringarna

import "fmt"

func Run1() {
	var string1, string2, string3, resultat string

	fmt.Printf("Sträng 1")
	fmt.Scanln(&string1)

	fmt.Printf("Sträng 2")
	fmt.Scanln(&string2)

	fmt.Printf("Sträng 3")
	fmt.Scanln(&string3)

	resultat = string1 + " " + string2 + " " + string3
	fmt.Println(resultat)

}
