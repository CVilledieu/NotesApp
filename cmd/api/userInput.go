package notes

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Nav() {
	fmt.Println("Would you like to:")
	fmt.Println("Enter 1 to start a new note")
	fmt.Println("Enter 2 to search for a note")
	fmt.Println("Enter 3 to view all notes")
	fmt.Println("Enter 4 to exit")
	for {
		fmt.Print("Enter your choice: ")
		text := getInputText()
		switch text {
		case "1":
			NewNote()
			return
		case "2":
			SearchNote()
			return
		case "3":
			ViewAll()
			return
		case "4":
			Exit()
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func ViewAll() {
	ClearScreen()
	fmt.Println("Viewing notes...")
	fmt.Print("\n")
	files, err := os.ReadDir("public/notes")
	if errIsNotNil(err) {
		return
	}
	for _, file := range files {
		fName := strings.TrimSuffix(file.Name(), ".txt")
		if strings.HasPrefix(fName, "_") {
			continue
		} else {
			fmt.Println(fName)
		}

	}
	viewAllMenuFollowUp()
}

func viewNote(path string) {
	ClearScreen()
	fmt.Println("Viewing note...")
	dividers()
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
			if err == io.EOF {
				viewNoteFollowUp()
				break
			} else {
				fmt.Println("Error reading file: ", err)
				return
			}

		}
		fmt.Println(string(data[:n]))
		if n == 0 {
			break
		}
	}
}
