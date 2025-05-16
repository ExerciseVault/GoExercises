package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//create a reader to read user input
	reader := bufio.NewReader(os.Stdin)

	//length
	fmt.Print("Enter the length of a rectangle: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, _ := strconv.ParseFloat(lengthInput, 64)

	//width
	fmt.Print("Enter the width of a rectangle: ")
	widthInput, _ := reader.ReadString('\n')
	widthInput = strings.TrimSpace(widthInput)
	width, _ := strconv.ParseFloat(widthInput, 64)

	//calculating the area
	area := length * width
	fmt.Println("Area of a rectangle ", area)

}
