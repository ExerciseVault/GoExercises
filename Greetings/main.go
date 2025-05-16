package main

import (
	"bufio"
    "fmt"
    "os"
    "strings"
)

func main(){

	//create a reader to read user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name : ")

	// Read the input until the newline character
	name, _ := reader.ReadString('\n')

	// Trim the newline character from the input
	name = strings.TrimSpace(name)

	fmt.Println("Hello, " + name + "! Nice to meet you.")
}