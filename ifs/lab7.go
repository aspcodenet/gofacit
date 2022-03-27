package ifs

import (
	"fmt"
	"strings"
)

func Run7() {
	var country string

	fmt.Print("Ange land:")
	fmt.Scanln(&country)

	country = strings.ToLower(country)
	if country == "sverige" || country == "danmark" || country == "norge" {
		fmt.Println("Du bor i Scandinavien")
	} else {
		fmt.Println("Du bor inte i Scandinavien")
	}

}

// Yepp... no generics etc.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Run7b() {

	var country string
	countries := []string{"sverige", "danmark", "norge"}

	fmt.Print("Ange land:")
	fmt.Scanln(&country)

	country = strings.ToLower(country)

	if contains(countries, country) {
		fmt.Println("Du bor i Scandinavien")
	} else {
		fmt.Println("Du bor inte i Scandinavien")
	}

}
