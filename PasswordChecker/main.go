package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a password to check its strength: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	var hasUpper, hasLower, hasNumber bool

	for _, character := range password {
		switch {
		case unicode.IsUpper(character):
			hasUpper = true
		case unicode.IsLower(character):
			hasLower = true
		case unicode.IsDigit(character):
			hasNumber = true
		}
	}

	length := len(password)

	fmt.Println("Password Strength Analysis:")
	if length >= 8 {
		fmt.Println("Length: OK")
	} else {
		fmt.Println("Length: Too short (min 8 characters)")
	}

	if hasUpper {
		fmt.Println("Contains uppercase letter: Yes")
	} else {
		fmt.Println("Contains uppercase letter: No")
	}

	if hasLower {
		fmt.Println("Contains lowercase letter: Yes")
	} else {
		fmt.Println("Contains lowercase letter: No")
	}

	if hasNumber {
		fmt.Println("Contains number: Yes")
	} else {
		fmt.Println("Contains number: No")
	}

	if length >= 8 && hasUpper && hasLower && hasNumber {
		fmt.Println("\nPassword strength: Strong ")
	} else {
		fmt.Println("\nPassword strength: Weak")
	}
}
