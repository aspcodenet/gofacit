package funktioner

import (
	"errors"
	"fmt"
)

func Run4() {
	tax, err := CalculateTaxesOnSalary(15800.0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Skatten blir %f\n", tax)
	}

}

func CalculateTaxesOnSalary(salary float64) (tax float64, err error) {
	if salary < 0 {
		return 0, errors.New("Inte negativ lön väl?")
	}
	if salary < 15000 {
		return salary * 0.12, nil
	}
	if salary < 30000 {
		return salary * 0.28, nil
	}
	return salary * 0.33, nil
}
