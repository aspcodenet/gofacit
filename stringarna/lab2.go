package stringarna

import (
	"fmt"
	"strings"
)

func Run2() {
	string1 := "Hello, world"

	pos1 := strings.Index(string1, "e")
	pos2 := strings.LastIndex(string1, "o")

	fmt.Printf("Positioner:%d ioch %d\n", pos1, pos2)

}
