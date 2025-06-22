package io

import (
	"fmt"
	"hangman/game"
	"strings"
)

var hangmanStates = [6]string{
	`
 +---+
 |   |
 O   |
     |
     |
     |
======
`,
	`
 +---+
 |   |
 O   |
 |   |
     |
     |
======
`,
	`
 +---+
 |   |
 O   |
/|   |
     |
     |
======
`,
	`
 +---+
 |   |
 O   |
/|\  |
     |
     |
======
`,
	`
 +---+
 |   |
 O   |
/|\  |
/    |
     |
======
`,
	`
 +---+
 |   |
 O   |
/|\  |
/ \  |
     |
======
`,
}

type OutPut struct{}

func NewOutput() *OutPut {
	return &OutPut{}
}

func (o *OutPut) RenderGame(g *game.Hangman) {
	fmt.Println("\nWord:")
	chars := make([]rune, len(g.Word))
	for i, ch := range g.Word {
		if g.Guessed[i] {
			chars[i] = ch
		} else {
			chars[i] = '_'
		}
	}

	fmt.Println("\n" + findTableWithLetters(chars))
	fmt.Println("\nAttempts left:", g.AttemptsLeft)
	fmt.Println("Tried letters:", keys(g.TriedLetters))
	renderHangman(len(hangmanStates) - g.AttemptsLeft - 1)

}
func (o *OutPut) RenderFeedback(msg string) {
	fmt.Println(msg)
}

func (o *OutPut) RenderResult(g *game.Hangman) {
	if g.IsWon() {
		fmt.Println("\nCongratulations! You`ve won")
	} else {
		fmt.Println("\nGame over! The word was: ", g.Word)
	}
}

func keys(m map[rune]bool) string {
	var sb strings.Builder
	for k := range m {
		sb.WriteRune(k)
		sb.WriteString(" ")
	}
	return sb.String()
}

func findTableWithLetters(arr []rune) string {
	builder := strings.Builder{}

	for i := 0; i < len(arr); i++ {

		builder.WriteString("+---")
	}
	builder.WriteString("+\n")

	for i := 0; i < len(arr); i++ {
		builder.WriteString("| " + string(arr[i]) + " ")
	}

	builder.WriteString("|\n")

	for i := 0; i < len(arr); i++ {

		builder.WriteString("+---")
	}

	builder.WriteString("+\n")

	return builder.String()
}

func renderHangman(index int) {
	if index >= 0 && index < len(hangmanStates) {
		fmt.Println(hangmanStates[index])
	} else {
		fmt.Println("(not error - no image)")
	}
}
func GetHangmanStatesCount() int {
	return len(hangmanStates)
}
