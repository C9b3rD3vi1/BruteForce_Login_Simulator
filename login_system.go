// login_system.go
package main

import (
	//"fmt"
	"sync"
	"github.com/fatih/color"
)


var (
	validUsername   = "admin"
	validPassword   = "securepassword"
	loginAttempts   = make(map[string]int)
	mutex           sync.Mutex
	maxLoginAttempts = 7 // max number of login attempts

    red = color.New(color.FgRed).Add(color.Underline)
	green = color.New(color.FgGreen)
)

// Login function to authenticate user with username and password
func Login(username, password string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if loginAttempts[username] >= maxLoginAttempts {
		red.Println("Account locked due to too many failed attempts.")
		return false
	}

	if username == validUsername && password == validPassword {
		green.Println("Login successful!")
		loginAttempts[username] = 0 // Reset attempts on success
		return true
	}
	
	loginAttempts[username]++
	remaingAttempts := maxLoginAttempts - loginAttempts[username]
	
	if remaingAttempts > 0 {
        red.Printf("Invalid credentials. You have %d remaining attempts.\n", remaingAttempts)
    } else {
        red.Println("Invalid credentials. Account locked due to too many failed attempts.")
    }

	return false
}
