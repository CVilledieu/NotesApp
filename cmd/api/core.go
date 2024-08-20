package notes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewNote() {
	if !safetyCheck() {
		return
	}
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
	_, err = file.WriteString("\n")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
}

func ViewNote(path string) {
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

func saveNote() {
	fmt.Println("Saving note...")
	tempFile, err := os.Open("public/notes/temp.txt")
	if errIsNil(err) {
		return
	}

	reader := bufio.NewReader(tempFile)
	title, err := reader.ReadString('\n')
	if errIsNil(err) {
		return
	}
	title = strings.TrimSuffix(title, "\n")
	newFile, err := os.OpenFile("public/notes/"+title+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errIsNil(err) {
		return
	}
	defer newFile.Close()
	defer tempFile.Close()
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if errIsNil(err) {
			return
		}
		_, err = newFile.WriteString(data)
		if errIsNil(err) {
			return
		}
	}
	tempFile.Close()
	err = os.Remove("public/notes/temp.txt")
	if errIsNil(err) {
		return
	}
}
