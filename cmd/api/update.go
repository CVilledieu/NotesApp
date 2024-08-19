package notes

import (
	"fmt"
	"os"
)

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
}
