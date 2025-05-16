package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Convert Celsius to Fahrenheit
func cToF(c float64) float64 {
	return (c * 9 / 5) + 32
}

// Convert Fahrenheit to Celsius
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for temperature
	fmt.Print("Enter the temperature value: ")
	tempInput, _ := reader.ReadString('\n')
	tempInput = strings.TrimSpace(tempInput)
	temp, _ := strconv.ParseFloat(tempInput, 64)

	// Ask for conversion type
	fmt.Print("Convert to (C)elsius or (F)ahrenheit? Enter C or F: ")
	scale, _ := reader.ReadString('\n')
	scale = strings.TrimSpace(strings.ToUpper(scale))

	// Decide and convert
	if scale == "C" {
		result := fToC(temp)
		fmt.Printf("%.2f°F is %.2f°C\n", temp, result)
	} else if scale == "F" {
		result := cToF(temp)
		fmt.Printf("%.2f°C is %.2f°F\n", temp, result)
	} else {
		fmt.Println("Invalid option. Please enter 'C' or 'F'.")
	}
}
