package main

import (
	"fmt"
	"main/fors"
	"main/ifs"
	"main/start"
	"main/strukturer"
	"math"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func main() {
	strukturer.Run3()
	fors.Run6()
	ifs.Run5()

	var arr1 = []int{1, 2, 3}

	fmt.Println(arr1)

	ifs.Run1()
	start.Run7()
	// t1 := 0.1
	// t2 := 0.2
	// t3 := 0.3

	// fmt.Printf("%9.9f", t1)
	// fmt.Printf("%9.9f", t2)
	// fmt.Printf("%9.9f", t3)

	// if t1+t2 == t3 {
	// 	fmt.Println("Jepp")
	// }

	// if almostEqual(t1+t2, t3) {
	// 	fmt.Println("Jepp")
	// }

	// start.Run6()
	// start.Run5()
	// start.Run4()
	// start.Run1()
	// start.Run2()
	// start.Run3()

}
