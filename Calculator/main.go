package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Functions for operations
func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// Get first number
	fmt.Print("Enter the first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, _ := strconv.ParseFloat(input1, 64)

	// Get second number
	fmt.Print("Enter the second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, _ := strconv.ParseFloat(input2, 64)

	// Get the operation
	fmt.Print("Enter operation (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Decide which function to call
	var result float64

	if op == "+" {
		result = add(num1, num2)
	} else if op == "-" {
		result = subtract(num1, num2)
	} else if op == "*" {
		result = multiply(num1, num2)
	} else if op == "/" {
		result = divide(num1, num2)
	} else {
		fmt.Println("Invalid operation.")
		return
	}

	fmt.Println("Result = ", result)

}
