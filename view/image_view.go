package view

import (
	"fmt"
	"strings"
)

var hangmanStates = []string{
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

func FindImageHangman(numberImage int) {
	if numberImage >= 0 && numberImage < len(hangmanStates) {
		fmt.Println(hangmanStates[numberImage])
	} else {
		fmt.Println("Нет изображения для такой ошибки")
	}
}

func FindTableWithLetters(arr []rune) string {
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

func GetHangmanStatesCount() int {
	return len(hangmanStates)
}
