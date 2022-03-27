package structs

import (
	"fmt"
	"time"
)

type Person struct {
	Birthdate  time.Time
	Namn       string
	GatuAdress string
	PostNummer int
	Postort    string
}

func (p1 *Person) MoveInTo(p *Person) {
	p1.GatuAdress = p.GatuAdress
	p1.PostNummer = p.PostNummer
	p1.Postort = p.Postort
}

func Run2() {
	p := Person{
		Birthdate:  time.Date(1972, time.August, 3, 0, 0, 0, 0, time.UTC),
		Namn:       "Stefan",
		GatuAdress: "Testgatan 11",
		PostNummer: 11122,
		Postort:    "Teststad",
	}
	p2 := Person{
		Birthdate:  time.Date(1973, time.March, 5, 0, 0, 0, 0, time.UTC),
		Namn:       "Kerstin",
		GatuAdress: "Hejvägen 1",
		PostNummer: 22233,
		Postort:    "Ettstad",
	}

	p.MoveInTo(&p2)
	fmt.Println(p)
	fmt.Println(p2)
}
