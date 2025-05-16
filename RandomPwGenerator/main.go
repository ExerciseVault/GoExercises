package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter desired password length: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	var length int
	fmt.Sscanf(lengthInput, "%d", &length)

	// Character sets to choose from
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()-_=+[]{}<>?/"

	// Ask user which character types to include
	fmt.Print("Include lowercase letters? (y/n): ")
	includeLower, _ := reader.ReadString('\n')
	includeLower = strings.TrimSpace(includeLower)

	fmt.Print("Include uppercase letters? (y/n): ")
	includeUpper, _ := reader.ReadString('\n')
	includeUpper = strings.TrimSpace(includeUpper)

	fmt.Print("Include numbers? (y/n): ")
	includeNumbers, _ := reader.ReadString('\n')
	includeNumbers = strings.TrimSpace(includeNumbers)

	fmt.Print("Include special characters? (y/n): ")
	includeSpecial, _ := reader.ReadString('\n')
	includeSpecial = strings.TrimSpace(includeSpecial)

	// Build the character pool
	charPool := ""
	if strings.ToLower(includeLower) == "y" {
		charPool += lower
	}
	if strings.ToLower(includeUpper) == "y" {
		charPool += upper
	}
	if strings.ToLower(includeNumbers) == "y" {
		charPool += numbers
	}
	if strings.ToLower(includeSpecial) == "y" {
		charPool += special
	}

	if len(charPool) == 0 {
		fmt.Println("No character types selected. Cannot generate password.")
		return
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate password
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charPool[rand.Intn(len(charPool))]
	}

	fmt.Println("Generated password:", string(password))
}
