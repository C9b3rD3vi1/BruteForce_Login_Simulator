package main
import "fmt"


func main() {

    // Brute force attack
	fmt.Println("Entering password guessing mode...")
	bruteForceAttack()
	
	for {
		var username, password string
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		if login(username, password) {
			break
		}
	}
}
