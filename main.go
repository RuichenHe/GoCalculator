package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		CurrentTime := time.Now()
		fmt.Println("Current time:", CurrentTime)

		aFloat1 := ParseInputValue("Value 1:")
		aFloat2 := ParseInputValue("Value 2:")
		InputOperation := ParseOperation("Operation:")
		FinalResult := CalculatorMain(aFloat1, aFloat2, InputOperation)
		fmt.Printf("%v %v %v = %v\n", aFloat1, InputOperation, aFloat2, FinalResult)
		if !ParseContinue("Continue? (y/n)") {
			break
		}
	}
}
func addValues(value1 float64, value2 float64) float64 {
	return math.Round((value1+value2)*100) / 100
}
func subtractValues(value1 float64, value2 float64) float64 {
	return math.Round((value1-value2)*100) / 100
}
func multiplyValues(value1 float64, value2 float64) float64 {
	return math.Round((value1*value2)*100) / 100
}
func divideValues(value1 float64, value2 float64) float64 {
	return math.Round((value1/value2)*100) / 100
}
func CalculatorMain(value1 float64, value2 float64, OperationChoice string) float64 {
	var CalResult float64
	switch OperationChoice {
	case "+":
		CalResult = addValues(value1, value2)
	case "-":
		CalResult = subtractValues(value1, value2)
	case "*":
		CalResult = multiplyValues(value1, value2)
	case "%":
		CalResult = divideValues(value1, value2)
	}
	return CalResult
}
func ParseInputValue(label string) float64 {
	fmt.Println(label)
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		panic("Invalid input")
	}
	return aFloat
}
func ParseOperation(label string) string {
	fmt.Println(label)
	fmt.Print("+ - * %:")
	numInput, _ := reader.ReadString('\n')
	InputOperation := strings.TrimSpace(numInput)
	return InputOperation
}
func ParseContinue(label string) bool {
	fmt.Println(label)
	fmt.Print("y/n:")
	numInput, _ := reader.ReadString('\n')
	InputOperation := strings.TrimSpace(numInput)
	switch InputOperation {
	case "Y":
		return true
	case "y":
		return true
	default:
		return false
	}
}
