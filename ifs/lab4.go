package ifs

import "fmt"

func Run4() {
	var tal1, tal2, tal3, largest int

	fmt.Printf("Ange tal 1:")
	fmt.Scanln(&tal1)
	fmt.Printf("Ange tal 2:")
	fmt.Scanln(&tal2)
	fmt.Printf("Ange tal 3:")
	fmt.Scanln(&tal3)

	if tal1 > tal2 && tal1 > tal3 {
		largest = tal1
	} else if tal2 > tal1 && tal2 > tal3 {
		largest = tal2
	} else {
		largest = tal3
	}

	fmt.Printf("Största talet är %d\n", largest)

}
