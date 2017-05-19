package server_test

import (
	"testing"
    "fmt"

    "github.com/LeoCBS/garden/pointers"
)

func TestSumPointer(t *testing.T) {
	calculator := pointers.NewCalculatorPointer(100)
    calculator.Sum(100)
    value := calculator.GetMemory()
    fmt.Printf("value " , value)
    if value != 200{
        t.Fatal("calculator don't sum value")
    }
}

func TestSumPerValue(t *testing.T) {
	calculator := pointers.NewCalculatorValue(100)
    calculator.SumNoPointer(100)
    value := calculator.GetMemoryNoPointer()
    fmt.Printf("value " , value)
    if value != 100{
        t.Fatal("calculator sum value")
    }
}