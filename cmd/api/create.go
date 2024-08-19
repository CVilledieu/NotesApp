package notes

import (
	"fmt"
	"os"
)

func NewNote() {

	fmt.Println("Enter the title of your note: ")

	title := getInputText()

	fmt.Println("Please enter !Exit to save and exit the note.")
	fmt.Println("Enter the content of your note: ")

	writeNote(title)
	for {
		content := getInputText()
		if content == "!Exit" {
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
