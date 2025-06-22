package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Input struct {
	reader *bufio.Reader
}

func NewInput() *Input {
	return &Input{reader: bufio.NewReader(os.Stdin)}
}

func (i *Input) GetLetter() rune {
	for {
		fmt.Println("Enter a letter: ")
		input, err := i.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		input = strings.TrimSpace(input)
		if len([]rune(input)) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		ch := []rune(input)[0]
		if !unicode.Is(unicode.Latin, ch) {
			fmt.Println("Only Latin letters are allowed.")
			continue
		}

		return unicode.ToLower(ch)
	}
}

func (i *Input) ConfirmStart() bool {
	const startGame = 'y'
	fmt.Println("Do you want to play hangman? 'y' - yes; any letter - no;")
	ch := i.GetLetter()
	return ch == startGame
}
