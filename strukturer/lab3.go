package strukturer

import (
	"fmt"
	"math/rand"
)

func Run3() {
	var konton map[string]int = make(map[string]int, 1)

	for {
		fmt.Println("**********************")
		fmt.Println("1. Skapa konto")
		fmt.Println("2. Ta bort konto")
		fmt.Println("3. Lista alla kontonummer")
		fmt.Println("4. Lista totalsaldo")
		fmt.Println("5. Lista alla kontonummer och saldo")

		var sel string
		fmt.Printf("Val:")
		fmt.Scanln(&sel)

		if sel == "1" {
			var konto string
			fmt.Println("Ange nytt kontonummer:")
			fmt.Scanln(&konto)
			konton[konto] = rand.Intn(1000)
		} else if sel == "2" {
			var konto string
			fmt.Println("Ange kontonummer att ta bort:")
			fmt.Scanln(&konto)
			delete(konton, konto)
		} else if sel == "3" {
			fmt.Println("Alla kontonummer")
			for konto := range konton {
				fmt.Println(konto)
			}
		} else if sel == "4" {
			fmt.Println("Totalsaldo")
			summa := 0
			for konto := range konton {
				summa += konton[konto]
			}
			fmt.Printf("Saldo:%d\n", summa)
		} else if sel == "5" {
			fmt.Println("Alla konton och saldon")
			for konto := range konton {
				fmt.Printf("%s:%d\n", konto, konton[konto])
			}
		}

	}
}
