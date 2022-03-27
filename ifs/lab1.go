package ifs

import "fmt"

func Run1() {
	var tal int
	fmt.Printf("Ange ett tal:")
	fmt.Scanln(&tal)

	if tal > 10 {
		fmt.Printf("Talet är större än 10")
	} else if tal < 10 {
		fmt.Printf("Talet är mindre än 10")
	}

}
