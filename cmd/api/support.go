package notes

import (
	"bufio"
	"fmt"
	"os"
)

func errIsNotNil(target interface{}) bool {
	if target != nil {
		dividers()
		fmt.Println("Error: ", target)
		return true
	}
	return false
}

func getInputText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func UnsavedWorkCheck() bool {
	if _, err := os.Stat("public/notes/_temp.txt"); os.IsNotExist(err) {
		return false
	}
	// If there is a temp file, check if the user wants to save it
	warning()
	return false
}

func warning() func() {
	dividers()
	fmt.Println("Warning: You have an unsaved note.")
	fmt.Println("Would you like to view the note? (Y/N)")
	save := getInputText()
	if save == "Y" || save == "y" {

		return handleUnsavedWork()
	} else if save == "N" || save == "n" {
		deleteNote("_temp")
		return nil
	} else {
		fmt.Println("Invalid input. Please try again.")

		return warning()
	}

}

func handleUnsavedWork() func() {
	ClearScreen()
	viewNote("public/notes/_temp.txt")
	dividers()
	fmt.Println("Save the note? (Y/N)")
	fmt.Println("Edit the note? (E)")
	view := getInputText()

	if view == "Y" || view == "y" {
		saveNote()
	} else if view == "N" || view == "n" {
		return nil
	} else if view == "E" || view == "e" {
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
				return nil
			}
			writeNote(content)
		}
	} else {
		fmt.Println("Invalid input. Please try again.")
		return handleUnsavedWork()
	}
	return nil
}

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func ClearScreen() {
	n := 20
	for i := 0; i < n; i++ {
		fmt.Print("\n")
	}
}

func dividers() {
	fmt.Println("------------------------------------------------")
}

func deleteNote(input string) {
	input = "public/notes/" + input + ".txt"
	err := os.Remove(input)
	if errIsNotNil(err) {
		return
	}
	fmt.Println("Note deleted.")
}
