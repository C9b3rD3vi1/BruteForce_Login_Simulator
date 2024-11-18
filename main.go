package main

import "fmt"

import (
    "github.com/fatih/color"
	"time"

)


func slowPrint(text string, delay time.Duration) {
    for _, char := range text {
        fmt.Print(string(char))  // Print one character at a time
        time.Sleep(delay)        // Add a delay between each character
    }
    fmt.Println()  // Print a newline after the slow print
}



func main() {

	 // Create color objects
	 red := color.New(color.FgRed).Add(color.Underline)
	 //green := color.New(color.FgGreen)
	 //blue := color.New(color.FgBlue)


    red.Println("Choose mode:")
    fmt.Println("1: Run brute force attack")
    fmt.Println("2: Interactive login")
    var choice int
	// Scan user input for choice
	fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    if choice == 1 {
        slowPrint("Entering password guessing mode...", 100*time.Millisecond)  // Slow print
		// call Brute force attack
        BruteForceAttack()

    } else if choice == 2 {
        slowPrint("Entering interactive login mode...",100*time.Millisecond)
		// Run interactive login loop
        for {
            var username, password string
            fmt.Print("Enter username: ")
            fmt.Scan(&username)
            fmt.Print("Enter password: ")
            fmt.Scan(&password)
            if Login(username, password) {
                break
            }
        }
    } else {
        red.Println("Invalid choice.")
    }
}
