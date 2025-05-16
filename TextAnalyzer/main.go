package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get user input
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Count characters (including spaces)
	charCount := len([]rune(input)) // handles Unicode

	// Count words
	words := strings.Fields(input)
	wordCount := len(words)


	// Display results
	fmt.Println("Text analysis:")
	fmt.Println("Total characters:", charCount)
	fmt.Println("Total words:", wordCount)
}
