// login_system.go
package main

import (
	"fmt"
	"sync"
)


var (
	validUsername   = "admin"
	validPassword   = "securepassword"
	loginAttempts   = make(map[string]int)
	mutex           sync.Mutex
	maxLoginAttempts = 5
)



func login(username, password string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if loginAttempts[username] >= maxLoginAttempts {
		fmt.Println("Account locked due to too many failed attempts.")
		return false
	}

	if username == validUsername && password == validPassword {
		fmt.Println("Login successful!")
		return true
	}

	loginAttempts[username]++
	fmt.Println("Invalid credentials.")
	return false
}
