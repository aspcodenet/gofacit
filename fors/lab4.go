package fors

import (
	"fmt"
)

func Run4() {
	var tal1, summa int

	for i := 0; i <= 10; i++ {
		fmt.Print("Ange tal:")
		fmt.Scanln(&tal1)

		summa += tal1

	}
	fmt.Printf("%d\n", summa)

}
