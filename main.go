// main.go
package main

import (
	"fmt"
	"time"
)

// artemis is the main function displaying the initial menu
func artemis() {
	SetConsoleTitle()
	fmt.Println("Logging in on:", time.Now().Format("2006-01-02"), ", what a beautiful day!")

	for {
		var whatToDo string
		fmt.Print(`What do you want to do today with this program?
1. Enter the menu
2. Get help
3. About
x. Exit the program
b. Restart the program
[>] `)

		fmt.Scanln(&whatToDo) // Read user input

		switch whatToDo {
		case "1":
			clearConsole()
			showMenuWithExecution()
		case "2":
			clearConsole()
			getHelp()
		case "3":
			clearConsole()
			showAbout()
		case "x":
			fmt.Println("Exiting the program. Goodbye!")
			wait(1)
			return
		case "b":
			fmt.Println("Restarting the program. Goodbye!")
			wait(1)
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}

func main() {
	artemis()
}
