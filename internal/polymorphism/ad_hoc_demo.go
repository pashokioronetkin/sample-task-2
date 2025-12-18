package polymorphism

import "fmt"

// Calculator интерфейс для демонстрации ad hoc полиморфизма
type Calculator interface {
	Add(a, b interface{}) interface{}
}

// IntCalculator для целых чисел
type IntCalculator struct{}

func (c IntCalculator) Add(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

// FloatCalculator для дробных чисел
type FloatCalculator struct{}

func (c FloatCalculator) Add(a, b interface{}) interface{} {
	return a.(float64) + b.(float64)
}

// Demo демонстрирует полиморфизм
func Demo() {
	fmt.Println("=== Ad hoc полиморфизм ===")

	var calc Calculator

	calc = IntCalculator{}
	fmt.Printf("Int: 5 + 3 = %v\n", calc.Add(5, 3))

	calc = FloatCalculator{}
	fmt.Printf("Float: 2.5 + 3.7 = %v\n", calc.Add(2.5, 3.7))
}
