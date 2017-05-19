package server

import "fmt"

type Calculator struct {
	Memory int
}

func (calculator *Calculator) Sum(value int) {
	fmt.Println("memory pointer: %s", &calculator.Memory)
	calculator.Memory = calculator.Memory + 100
}

func (calculator *Calculator) GetMemory() int {
	fmt.Println("memory pointer: %s", &calculator.Memory)
	return calculator.Memory
}

func NewCalculatorPointer(initValue int) *Calculator {
	return &Calculator{
		Memory: initValue,
	}
}

func NewCalculatorValue(initValue int) Calculator {
	return Calculator{
		Memory: initValue,
	}
}

func (calculator Calculator) SumNoPointer(value int) {
	fmt.Println("memory pointer: %s", &calculator.Memory)
	calculator.Memory = calculator.Memory + 100
}

func (calculator Calculator) GetMemoryNoPointer() int {
	fmt.Println("memory pointer: %s", &calculator.Memory)
	return calculator.Memory
}
