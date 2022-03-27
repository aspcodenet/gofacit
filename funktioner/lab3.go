package funktioner

import "fmt"

func Run3() {
	ord := []string{"Kalle", "Elisabeth", "Bo"}

	langst := HittaLangstaOrder(ord)
	fmt.Println(langst)
}

func HittaLangstaOrder(ord []string) (longest string) {
	longest = ord[0]
	for _, val := range ord {
		if len(val) > len(longest) {
			longest = val
		}
	}
}
