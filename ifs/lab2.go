package ifs

import "fmt"

func Run2() {
	fmt.Printf("Hur många paket mjölk finns kvar?")
	var milk int
	fmt.Scanln(&milk)

	if milk < 10 {
		fmt.Println("Beställ 30 paket")
	} else if milk >= 10 && milk <= 20 {
		fmt.Println("Beställ 20 paket")
	} else {
		fmt.Println("Du behöver inte beställa nån mjölk")
	}
}
