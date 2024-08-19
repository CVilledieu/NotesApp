package main

import (
	notes "Notes/cmd/api"
	welcome "Notes/cmd/banner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome.Banner()
	Nav()
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
	fmt.Println("Enter 3 to exit")
	for {
		fmt.Print("Enter your choice: ")
		text := getInputText()
		switch text {
		case "1":
			notes.NewNote()
		case "2":
			searchNote()
		case "3":
			notes.Exit()
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func searchNote() {
	fmt.Println("Searching for notes...")
	fmt.Println("Enter ViewAll to view all notes.")
	fmt.Println("Enter the title of the note you would like to search for: ")
}

func listNotes() {
	fmt.Println("Here are your notes:")
}

func editNote() {
	fmt.Println("Enter the title of the note you would like to edit: ")
}

func deleteNote() {
	fmt.Println("Enter the title of the note you would like to delete: ")
}
