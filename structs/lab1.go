package structs

import "fmt"

type MatrattType int8

const (
	Vegetarisk MatrattType = iota
	Vegansk
	Kott
)

type Matratt struct {
	Namn     string
	Price    float64
	Typ      MatrattType
	Kalorier int
}

func Run1() {
	matratter := make([]Matratt, 0, 10)

	matratter = append(matratter, Matratt{Namn: "Pannkakor", Price: 12, Typ: Vegansk, Kalorier: 20})
	matratter = append(matratter, Matratt{Namn: "Pizza", Price: 22, Typ: Kott, Kalorier: 30})

	for _, matratt := range matratter {
		fmt.Println(matratt)
	}

}
