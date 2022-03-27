package funktioner

import "fmt"

func Run1() {
	res := ConcatenateStrings("Hejsan", "Hoppsan")
	fmt.Println(res)
}

func ConcatenateStrings(s1, s2 string) string {
	return s1 + s2
}
