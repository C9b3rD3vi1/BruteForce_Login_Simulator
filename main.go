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
	red := color.New(color.FgRed)  // For error messages
	orange := color.New(color.FgHiMagenta) // For prompts

	red.Println("Choose mode:")
	fmt.Println("1: Run brute force attack")
	fmt.Println("2: Interactive login")
	var choice int

	// Scan user input for choice
	for {
		orange.Print("Enter your choice: ") // Prompt in orange
		fmt.Scan(&choice) // Read the user input

		// Check if the choice is valid (1 or 2)
		if choice == 1 {
			slowPrint("Entering password guessing mode...", 100*time.Millisecond) // Slow print
			// Call Brute Force Attack function
			BruteForceAttack()
			break // Exit the loop after a valid choice

		} else if choice == 2 {
			slowPrint("Entering interactive login mode...", 100*time.Millisecond) // Slow print
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
			break // Exit the loop after a valid choice
			
		} else {
			red.Println("Invalid choice. Please select a valid option.") // Error message in red
		}
	}
}