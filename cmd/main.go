package main

import (
	welcome "Notes/cmd/banner"
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome.Banner()
	Nav()

}

func Nav() {
	fmt.Println("Would you like to:")
	fmt.Println("Enter 1 to start a new note")
	fmt.Println("Enter 2 to search for a note")
	fmt.Println("Enter 3 to exit")
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your choice: ")
		text, _ := scanner.ReadString('\n')
		switch text {
		case "1":
			newNote(scanner)
		case "2":
			searchNote()
		case "3":
			Exit()
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func newNote(scanner *bufio.Reader) {
	fmt.Println("Enter the title of your note: ")
	//title, _ := scanner.ReadString('\n')
	fmt.Println("Please enter !Exit to save and exit the note.")
	fmt.Println("Enter the content of your note: ")
	for {
		content, _ := scanner.ReadString('\n')
		if content == "!Exit" {
			saveNote()
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

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func saveNote() {
	fmt.Println("Saving note...")
}
