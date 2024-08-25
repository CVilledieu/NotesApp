package notes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewNote() {
	ClearScreen()
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
	ClearScreen()
	fmt.Println("Searching for notes...")
	fmt.Println("Enter ViewAll to view all notes.")
	fmt.Println("Enter the title of the note you would like to search for: ")
}

func viewAllMenuFollowUp() {
	fmt.Print("\n")
	fmt.Println("Enter !Menu to return to the main menu.")
	fmt.Println("Enter !Exit to exit the program.")
	fmt.Println("Enter the name of the note you would like to view: ")
	input := getInputText()
	input = strings.ToLower(input)
	if input == "!menu" {
		ClearScreen()
		return
	} else if input == "!Exit" || input == "!exit" {
		Exit()
	} else {
		f, err := os.OpenFile("public/notes/"+input+".txt", os.O_RDONLY, 0644)
		f.Close()
		if errIsNotNil(err) {
			fmt.Println("Note not found.")
		} else {
			viewNote("public/notes/" + input + ".txt")

		}
	}
}

func viewNoteFollowUp() {
	dividers()
	fmt.Println("Enter !Menu to return to the main menu.")
	fmt.Println("Enter !Exit to exit the program.")
	fmt.Println("Enter !Edit to edit the note.")
	input := getInputText()
	input = strings.ToLower(input)
	if input == "!menu" {
		ClearScreen()
		return
	} else if input == "!exit" {
		Exit()
	} else if input == "!edit" {
		fmt.Println("Please enter !Exit to save and exit the note.")
		fmt.Println("Enter the content of your note: ")
		for {
			content := getInputText()
			if content == "!Exit" || content == "!exit" {
				fmt.Println("Save the note? (Y/N)")
				save := getInputText()
				if save == "Y" {
					saveNote()
				}
				return
			}
			writeNote(content)
		}
	} else {
		fmt.Println("Invalid input. Please try again.")
		viewNoteFollowUp()
	}

}

func saveNote() {
	fmt.Println("Saving note...")
	tempFile, err := os.Open("public/notes/_temp.txt")
	if errIsNotNil(err) {
		return
	}

	reader := bufio.NewReader(tempFile)
	title, err := reader.ReadString('\n')
	if errIsNotNil(err) {
		return
	}
	title = strings.TrimSuffix(title, "\n")
	newFile, err := os.OpenFile("public/notes/"+title+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errIsNotNil(err) {
		return
	}
	defer newFile.Close()
	defer tempFile.Close()
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if errIsNotNil(err) {
			return
		}
		_, err = newFile.WriteString(data)
		if errIsNotNil(err) {
			return
		}
	}
	tempFile.Close()
	err = os.Remove("public/notes/_temp.txt")
	if errIsNotNil(err) {
		return
	}
}

func writeNote(content string) {
	file, err := os.OpenFile("public/notes/_temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	_, err = file.WriteString("\n")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
}
