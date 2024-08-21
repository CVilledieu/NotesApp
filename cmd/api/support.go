package notes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func errIsNotNil(target interface{}) bool {
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

func UnsavedWork() bool {
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
		viewNote("public/notes/temp.txt") //
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
	save := getInputText()
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

func saveNote() {
	fmt.Println("Saving note...")
	tempFile, err := os.Open("public/notes/temp.txt")
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
	err = os.Remove("public/notes/temp.txt")
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

func viewNote(path string) {
	clearScreen()
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

func clearScreen() {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Print("\n")
	}
}
