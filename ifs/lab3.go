package ifs

import "fmt"

func Run3() {
	var tal1, tal2, largest int

	fmt.Printf("Ange tal 1:")
	fmt.Scanln(&tal1)
	fmt.Printf("Ange tal 2:")
	fmt.Scanln(&tal2)

	// largest = tal1 > tal2 ? tal1 : tal2  = NO TERNARY OP
	// https://go.dev/doc/faq#Does_Go_have_a_ternary_form
	if tal1 > tal2 {
		largest = tal1
	} else {
		largest = tal2
	}

	fmt.Printf("Största talet är %d\n", largest)

}
