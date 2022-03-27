package fors

import "fmt"

func Run2() {
	var tal1, tal2 int

	fmt.Print("Ange tal1:")
	fmt.Scanln(&tal1)
	fmt.Print("Ange tal2:")
	fmt.Scanln(&tal2)

	for i := tal1; i <= tal2; i++ {
		fmt.Println(i)
	}
}
