package stringarna

import (
	"fmt"
	"strings"
)

func Run4() {
	texten := "Detta är en sträng som du skall ändra"

	texten = strings.ReplaceAll(texten, " ", "*")

	antal := strings.Count(texten, "*")

	fmt.Printf("%s %d\n", texten, antal)

}
