package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	printWelcome()

}

func printWelcome() {
	fmt.Println("Welcome to my notes!")
}

func navStart() {
	fmt.Println("Would you like to:")
	fmt.Println("Enter 1 to start a new note")
	fmt.Println("Enter 2 for a list of current notes")
	fmt.Println("Enter 3 to edit a prexisting note")
	fmt.Println("Enter 4 to delete a preexisting note")
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your choice: ")
		text, _ := scanner.ReadString('\n')
		switch text {
		case "1":
			newNote()
		case "2":
			listNotes()
		case "3":
			editNote()
		case "4":
			deleteNote()
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func newNote() {
	fmt.Println("Enter the title of your note: ")
	scanner := bufio.NewReader(os.Stdin)
	title, _ := scanner.ReadString('\n')
	fmt.Println("Enter the content of your note: ")
	content, _ := scanner.ReadString('\n')
	fmt.Println("Note created!")
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
