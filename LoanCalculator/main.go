package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to calculate monthly loan payment
func calculatePayment(principal, annualRate float64, months int) float64 {
	r := annualRate / 12 / 100 // monthly interest rate
	n := float64(months)

	if r == 0 {
		return principal / n
	}

	return principal * r * pow(1+r, n) / (pow(1+r, n) - 1)
}

// Power function (since math.Pow only takes float64)
func pow(a, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get principal
	fmt.Print("Enter loan principal amount: ")
	principalInput, _ := reader.ReadString('\n')
	principalInput = strings.TrimSpace(principalInput)
	principal, _ := strconv.ParseFloat(principalInput, 64)

	// Get interest rate
	fmt.Print("Enter annual interest rate (e.g. 5 for 5%): ")
	rateInput, _ := reader.ReadString('\n')
	rateInput = strings.TrimSpace(rateInput)
	annualRate, _ := strconv.ParseFloat(rateInput, 64)

	// Get loan term
	fmt.Print("Enter loan term (in months): ")
	termInput, _ := reader.ReadString('\n')
	termInput = strings.TrimSpace(termInput)
	months, _ := strconv.Atoi(termInput)

	// Calculate monthly payment
	payment := calculatePayment(principal, annualRate, months)

	// Display result
	fmt.Println("Your monthly payment is:", payment)
}
