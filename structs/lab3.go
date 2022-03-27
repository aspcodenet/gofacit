package structs

import (
	"fmt"
	"strings"

	"github.com/cheynewallace/tabby"
)

func mainMenu() (sel int) {
	for {
		fmt.Println("********************")
		fmt.Println("1. Registrera")
		fmt.Println("2. Lista")
		fmt.Println("3. Ta bort")
		fmt.Println("4. Avsluta")

		fmt.Scanln(&sel)
		if sel >= 1 && sel <= 4 {
			return
		}

		fmt.Println("Ej valid val")

	}

}

type Dog struct {
	Name     string
	Breed    string
	Age      int
	WeightKg float64
}

func (d *Dog) TailLength() float32 {
	if strings.ToLower(d.Breed) == "tax" {
		return 3.7
	}
	return float32(d.Age) * float32(d.WeightKg) / 10.0

}

func Run3() {
	kennel := make([]Dog, 0, 10)

	for {
		val := mainMenu()
		switch val {
		case 1:
			d := register()
			kennel = append(kennel, *d)
		case 2:
			listAll(kennel)
		case 3:
			index := getIndexToRemove(kennel)
			if index != -1 {
				kennel = remove(kennel, index)
			} else {
				fmt.Println("Finns ingen sån hund")
			}

		case 4:
			return
		}
	}
}

func getIndexToRemove(kennel []Dog) (index int) {
	var namn string
	fmt.Print("Ange namn på hund att ta bort:")
	fmt.Scanln(&namn)
	for i, val := range kennel {
		if val.Name == namn {
			return i
		}
	}
	return -1
}

func remove(s []Dog, i int) []Dog {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func register() *Dog {
	dog := Dog{}
	fmt.Println("**** Ny ****")
	fmt.Print("Namn:")
	fmt.Scanln(&dog.Name)
	fmt.Print("Breed:")
	fmt.Scanln(&dog.Breed)
	fmt.Print("Age:")
	fmt.Scanln(&dog.Age)
	fmt.Print("Weight:")
	fmt.Scanln(&dog.WeightKg)

	return &dog
}

func listAll(dog []Dog) {
	var minTailLength int

	fmt.Print("Ange minsta svanslängd:")
	fmt.Scanln(&minTailLength)

	fmt.Println("**** Lista ****")
	t := tabby.New()
	t.AddHeader("Name", "Breed", "Age", "Weight", "Tail length")
	for _, val := range dog {
		if val.TailLength() >= float32(minTailLength) {
			t.AddLine(val.Name, val.Breed, val.Age, val.WeightKg, val.TailLength())
		}
	}
	t.Print()
}
