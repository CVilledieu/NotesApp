package notes

import (
	"bufio"
	"fmt"
	"os"
)

func errIsNil(target interface{}) bool {
	if target != nil {
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

func safetyCheck() bool {
	// Check if there is a temp file
	if !isThereATemp() {
		return true // If there is no temp file, continue on with the program
	}
	// If there is a temp file, check if the user wants to save it
	if !warning() {
		return true // If the user does not want to save the note, delete the temp file
	}
	fmt.Println("Would you like to view the note before saving? (Y/N)")
	view := getInputText()

	if view == "Y" || view == "y" {
		ViewNote("public/notes/temp.txt") //
		return false
	} else if view == "N" || view == "n" {
		saveNote()
		return true
	} else {
		fmt.Println("Invalid input. Please try again.")
		return safetyCheck()
	}
}

func isThereATemp() bool {
	if _, err := os.Stat("public/notes/temp.txt"); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func warning() bool {
	fmt.Println("Warning: You have an unsaved note.")
	fmt.Println("Would you like to save it? (Y/N)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	save := scanner.Text()
	if save == "Y" || save == "y" {
		return true
	} else if save == "N" || save == "n" {
		return false
	} else {
		fmt.Println("Invalid input. Please try again.")
		return warning()
	}

}

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}
