package ifs

import (
	"fmt"
	"strings"
)

func Run5() {
	var kategori string
	var price int

	fmt.Printf("Ange kategori (vuxen, pensionär, student)")
	fmt.Scanln(&kategori)

	kategori = strings.ToLower(kategori)

	if kategori == "vuxen" || kategori == "v" {
		price = 30
	} else if kategori == "pensionär" || kategori == "p" || kategori == "student" || kategori == "s" {
		price = 20
	}

	if price > 0 {
		fmt.Printf("Resan kostar %d kr\n", price)
	} else {
		fmt.Println("Felaktig kategori")
	}

}
