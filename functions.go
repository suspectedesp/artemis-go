package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)
const (
	Title = "Artemis-go"	
)

func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func restart() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command(os.Args[0])
	case "linux", "darwin":
		cmd = exec.Command("./" + os.Args[0])
	default:
		fmt.Println("Unsupported operating system:", runtime.GOOS)
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error restarting:", err)
		return
	}

	os.Exit(0)
}

func SetConsoleTitle() {
	switch runtime.GOOS {
	case "windows":
		exec.Command("cmd", "/c", "title", Title).Run()
	case "linux", "darwin":
		// Linux and macOS code
		fmt.Printf("\033]0;%s\007", Title)
	}
}

func wait(scnds int) {
	time.Sleep(time.Duration(scnds) * time.Second)
}