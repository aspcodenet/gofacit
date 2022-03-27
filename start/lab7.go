package start

import "fmt"

func Run7() {
	var fornamn string
	fmt.Printf("Vad herer du:")
	fmt.Scanln(&fornamn)

	timesFive := fornamn + fornamn + fornamn + fornamn + fornamn

	fmt.Println(timesFive)
}
