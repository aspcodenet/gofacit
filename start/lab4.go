package start

import "fmt"

func Run4() {
	var fornamn string
	var efternamn string

	fmt.Printf("Skriv in ditt förnamn:")
	fmt.Scanln(&fornamn)

	fmt.Printf("Skriv in ditt efternamn:")
	fmt.Scanln(&efternamn)

	fmt.Printf("Du heter: %s, %s\n", efternamn, fornamn)
}
