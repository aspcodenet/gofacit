package funktioner

import "fmt"

func Run2() {
	moms := CalculateVat(12.0)
	fmt.Println(moms)
}

func CalculateVat(f float64) float64 {
	return f * 0.2
}
