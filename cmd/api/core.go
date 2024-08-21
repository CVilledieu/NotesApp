package notes

import (
	"fmt"
	"os"
	"strings"
)

func NewNote() {
	clearScreen()
	fmt.Println("Enter the title of your note: ")

	title := getInputText()

	fmt.Println("Please enter !Exit to save and exit the note.")
	fmt.Println("Enter the content of your note: ")

	writeNote(title)
	for {
		content := getInputText()
		if content == "!Exit" || content == "!exit" {
			fmt.Println("Save the note? (Y/N)")
			save := getInputText()
			if save == "Y" {
				saveNote()
			}
			Exit()
		} else {
			writeNote(content)
		}
	}

}

func SearchNote() {
	clearScreen()
	fmt.Println("Searching for notes...")
	fmt.Println("Enter ViewAll to view all notes.")
	fmt.Println("Enter the title of the note you would like to search for: ")
}

func ViewAll() {
	clearScreen()
	fmt.Println("Viewing notes...")
	fmt.Print("\n")
	files, err := os.ReadDir("public/notes")
	if errIsNotNil(err) {
		return
	}
	for _, file := range files {
		fName := strings.TrimSuffix(file.Name(), ".txt")
		fmt.Println(fName)
	}
	viewAllMenuFollowUp()
}

func viewAllMenuFollowUp() {
	fmt.Print("\n")
	fmt.Println("Enter !Menu to return to the main menu.")
	fmt.Println("Enter !Exit to exit the program.")
	fmt.Println("Enter the name of the note you would like to view: ")
	input := getInputText()
	input = strings.ToLower(input)
	if input == "!menu" {
		return
	} else if input == "!Exit" || input == "!exit" {
		Exit()
	} else {
		f, err := os.OpenFile("../public/notes/"+input+".txt", os.O_RDONLY, 0644)
		f.Close()
		if errIsNotNil(err) {

			viewNote("../public/notes/" + input + ".txt")
		} else {
			fmt.Println("Note not found.")
		}
	}
}
