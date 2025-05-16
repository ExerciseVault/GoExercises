package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())
    secret := rand.Intn(10) + 1 // Random number between 1 and 10

    reader := bufio.NewReader(os.Stdin)
    var guess int

    fmt.Println("Guess the number (between 1 and 10):")


    for {
        fmt.Print("Enter your guess: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        guess, _ = strconv.Atoi(input) // convert input to integer

        if guess < secret {
            fmt.Println("Too low! Try again.")
        } else if guess > secret {
            fmt.Println("Too high! Try again.")
        } else {
            fmt.Println("Correct! You guessed the number!")
            break // exit the loop
        }
    }
}
