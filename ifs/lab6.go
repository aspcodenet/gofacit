package ifs

import "fmt"

func main() {
	var age int

	fmt.Print("Ange födelseår:")
	fmt.Scanln(&age)

	if age >= 1980 && age < 1980 {
		fmt.Println("Du är född på 1980-talet")
	} else if age < 2000 && age >= 1990 {
		fmt.Println("Du är född på 1990-talet")
	} else {
		fmt.Println("Du är inte född på 1980- eller 1990-talen")
	}

}
