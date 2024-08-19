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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		text := scanner.Bytes()
		switch string(text) {
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

func newNote(scanner *bufio.Scanner) {
	fmt.Println("Enter the title of your note: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Println("Please enter !Exit to save and exit the note.")
	fmt.Println("Enter the content of your note: ")
	writeNote(title)
	for {
		scanner.Scan()
		content := scanner.Text()
		if content == "!Exit" {
			fmt.Println("Save the note? (Y/N)")
			scanner.Scan()
			save := scanner.Text()
			if save == "Y" {
				saveNote()
			}
			Exit()
		} else {
			writeNote(content)
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

func viewNote(path string) {
	fmt.Println("Viewing note...")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()
	data := make([]byte, 100)
	for {
		n, err := file.Read(data)
		if err != nil {
			fmt.Println("Error reading file: ", err)
			return
		}
		fmt.Println(string(data[:n]))
		if n == 0 {
			break
		}
	}
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

func writeNote(content string) {
	file, err := os.OpenFile("public/notes/temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
}
