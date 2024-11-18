package main

import "fmt"

import (
    "github.com/fatih/color"
)

func main() {

	 // Create color objects
	 red := color.New(color.FgRed).Add(color.Underline)
	 green := color.New(color.FgGreen)
	 blue := color.New(color.FgBlue)


    fmt.Println("Choose mode:")
    fmt.Println("1: Run brute force attack")
    fmt.Println("2: Interactive login")
    var choice int
    fmt.Scan(&choice)

    if choice == 1 {
        fmt.Println("Entering password guessing mode...")
        BruteForceAttack()

    } else if choice == 2 {
        fmt.Println("Entering interactive login mode...")
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
        fmt.Println("Invalid choice.")
    }
}
