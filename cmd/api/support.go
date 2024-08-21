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

func UnsavedWorkCheck() bool {
	if _, err := os.Stat("public/notes/_temp.txt"); os.IsNotExist(err) {
		return false
	}
	// If there is a temp file, check if the user wants to save it
	warning()
	return false
}

func warning() func() {
	fmt.Println("Warning: You have an unsaved note.")
	fmt.Println("Would you like to view the note? (Y/N)")
	save := getInputText()
	if save == "Y" || save == "y" {

		return handleUnsavedWork()
	} else if save == "N" || save == "n" {
		return nil
	} else {
		fmt.Println("Invalid input. Please try again.")

		return warning()
	}

}

func handleUnsavedWork() func() {
	clearScreen()
	viewNote("public/notes/_temp.txt")

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
