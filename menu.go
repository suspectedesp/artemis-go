// go/funcs/funcs.go
package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"os/exec"
	"runtime"
)

const scriptDir = "./scripts/"

// getMenuOptions returns a list of Python files in the ./scripts/ directory
func getMenuOptions() ([]string, error) {
	files, err := ioutil.ReadDir(scriptDir)
	if err != nil {
		return nil, fmt.Errorf("error reading script directory: %w", err)
	}

	var options []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".py" {
			options = append(options, file.Name())
		}
	}

	return options, nil
}

// showMenu displays the available Python scripts in the ./scripts/ directory
func showMenu() {
	ascii := `   _____          __                 .__        
  /  _  \________/  |_  ____   _____ |__| ______
 /  /_\  \_  __ \   __\/ __ \ /     \|  |/  ___/
/    |    \  | \/|  | \  ___/|  Y Y  \  |\___ \ 
\____|__  /__|   |__|  \___  >__|_|  /__/____  >
        \/                 \/      \/        \/`
    fmt.Println(ascii)
    fmt.Println("github.com/vortexsys")
    fmt.Println("Press x to exit, b to go back")
	fmt.Println("Select a Python script:")
	options, err := getMenuOptions()
	if err != nil {
		fmt.Printf("Error getting menu options: %v\n", err)
		return
	}

	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

// getSelectedScriptPath returns the file path of the selected Python script
func getSelectedScriptPath(selectedOption int) (string, error) {
	options, err := getMenuOptions()
	if err != nil {
		return "", fmt.Errorf("error getting script path: %w", err)
	}

	if selectedOption > 0 && selectedOption <= len(options) {
		return filepath.Join(scriptDir, options[selectedOption-1]), nil
	}

	return "", fmt.Errorf("invalid option")
}

// executeSelectedScript executes the selected Python script
func executeSelectedScript(selectedOption int) error {
	scriptPath, err := getSelectedScriptPath(selectedOption)
	if err != nil {
		return fmt.Errorf("error getting script path: %w", err)
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// For Windows, use "cmd" to open a new console window and execute the Python script
		cmd = exec.Command("cmd", "/c", "start", "cmd", "/c", "title Python Script && python", scriptPath)
	default:
		// For other platforms, just execute the Python script directly
		cmd = exec.Command("python", scriptPath)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err = cmd.Start()
	clearConsole()
	return nil
}

// showMenuWithExecution displays the menu and executes the selected Python script
func showMenuWithExecution() {
	showMenu()

	var selectedOption int
	fmt.Print("[>] ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// Check if 'x' is pressed to exit
	if input == "x" {
		fmt.Println("Exiting program.")
		wait(1)
		os.Exit(0)
	}

	if input == "b" {
		fmt.Println("Restarting program.")
		wait(1)
		clearConsole()
		restart()
	}

	// Parse the selected option
	_, err = fmt.Sscanf(input, "%d", &selectedOption)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		return
	}

	err = executeSelectedScript(selectedOption)
	if err != nil {
		fmt.Printf("Error executing script: %v\n", err)
	}
}
func getHelp() {
    fmt.Println("Please report any bug at https://github.com/vortexsys/artemis-go/issues")
    fmt.Println("With what do you need help?")
	fmt.Println("1. How to use")
	fmt.Println("2. Purpose")
	fmt.Print("[>] ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	if input == "1" {
		fmt.Println("1. Upload your python file to ./scripts/")
		fmt.Println("2. Start this program")
		fmt.Println("3. Go to the menu again and select it")
		wait(4)
		clearConsole()
	}
	if input == "2" {
		fmt.Println("Basically, this is software able to run Python files that are thrown together into the extensions folder with (mostly) no problems.")
		fmt.Println("The purpose behind this is to unite multiple projects into one powerful tool")
		wait(6)
		clearConsole()
	}
	if input != "1" && input != "2" {
		fmt.Println("Invalid Input!")
		wait(1)
		clearConsole()
		showMenuWithExecution()
	}
}

func showAbout() {
    fmt.Println("This is the about information.")
    // ... (rest of the function)
}