package main

import "fmt"

func main() {
	// Uncomment the desired mode

	// Mode 1: Run the brute force simulation
 fmt.Println("Entering password guessing mode...")
	bruteForceAttack()

	// Mode 2: Interactive login system
	// Uncomment this block to test interactive login
	/*
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
	*/
}
