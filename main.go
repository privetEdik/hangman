package main

import (
	"hangman/io"
	"hangman/session"
)

func main() {
	input := io.NewInput()

	for input.ConfirmStart() {
		seance := session.NewGameSession()
		seance.Run()
	}
}
