package main

import (
	"hangman/io"
	"hangman/session"
)

func main() {
	input := io.NewInput()

	for input.ConfirmStart() {
		session := session.NewGameSession()
		session.Run()
	}
}
