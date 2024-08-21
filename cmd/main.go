package main

import (
	notes "Notes/cmd/api"
	welcome "Notes/cmd/banner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	notes.ClearScreen()
	welcome.Banner()
	for {
		Nav()
	}
}

func getInputText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Nav() {
	fmt.Println("Would you like to:")
	fmt.Println("Enter 1 to start a new note")
	fmt.Println("Enter 2 to search for a note")
	fmt.Println("Enter 3 to view all notes")
	fmt.Println("Enter 4 to exit")
	for {
		fmt.Print("Enter your choice: ")
		text := getInputText()
		switch text {
		case "1":
			notes.NewNote()
			return
		case "2":
			notes.SearchNote()
			return
		case "3":
			notes.ViewAll()
			return
		case "4":
			notes.Exit()
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
