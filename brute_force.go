// brute_force.go
package main

import "fmt"

var passwords = []string{"12345", "password", "admin123", "securepassword"}

func bruteForceAttack() {
	username := "admin"
	for _, password := range passwords {
		fmt.Printf("Trying %s / %s\n", username, password)
		if login(username, password) {
			fmt.Printf("Password found: %s\n", password)
			return
		}
	}
	fmt.Println("Password not found.")
}