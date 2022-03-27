package fors

import (
	"fmt"
)

func Run5() {
	var tal1 int
	fmt.Print("Ange tal:")
	fmt.Scanln(&tal1)

	for i := tal1; i > 0; i-- {
		fmt.Println(i)
	}

}
