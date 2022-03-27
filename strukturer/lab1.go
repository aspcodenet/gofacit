package strukturer

import "fmt"

func Run1() {
	var talen [4]int

	for index := 0; index < 4; index++ {
		fmt.Print("Ange tal")
		fmt.Scanln(&talen[index])
	}

	largest := talen[0]
	for index := 1; index < 4; index++ {
		if talen[index] > largest {
			largest = talen[index]
		}
	}

	fmt.Printf("Största är %d\n", largest)

}
