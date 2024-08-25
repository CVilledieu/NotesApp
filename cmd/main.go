package main

import (
	notes "Notes/cmd/api"
	welcome "Notes/cmd/banner"
)

func main() {
	notes.ClearScreen()
	welcome.Banner()
	for {
		notes.Nav()
	}
}
