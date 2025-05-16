package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Open the input file
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer inputFile.Close()

    // Initialize a scanner to read the file line by line
    scanner := bufio.NewScanner(inputFile)

    // Variables to store processing results
    lineCount := 0
    wordCount := 0
    specificWord := "Go"
    specificWordCount := 0

    // Read and process each line
    for scanner.Scan() {
        line := scanner.Text()
        lineCount++

        // Split the line into words
        words := strings.Fields(line)
        wordCount += len(words)

        // Count occurrences of the specific word (case-insensitive)
        for _, word := range words {
            if strings.EqualFold(word, specificWord) {
                specificWordCount++
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
        return
    }

    // Prepare the result string
    result := fmt.Sprintf("Total Lines: %d\nTotal Words: %d\nOccurrences of '%s': %d\n",
        lineCount, wordCount, specificWord, specificWordCount)

    // Write the result to the output file
    err = os.WriteFile("output.txt", []byte(result), 0644)
    if err != nil {
        fmt.Println("Error writing to output file:", err)
        return
    }

    fmt.Println("Processing complete. Results written to output.txt")
}
