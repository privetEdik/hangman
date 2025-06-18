package controller

import (
	"bufio"
	"fmt"
	"hangman/repository"
	"hangman/service"
	"hangman/view"
	"os"
	"strings"
)

var limitErrors = view.GetHangmanStatesCount()
var reader = bufio.NewReader(os.Stdin)

func StartGame() {

	var startGame string
	for {

		fmt.Println("Хотитет сыграть в игру 'Виселица'?")
		fmt.Println("Введите литеру 'g' для начала игры или что угодно для завершения и нажмите Enter")

		_, err := fmt.Scanln(&startGame)
		if err != nil {
			return
		}

		if strings.TrimSpace(startGame) != "g" {
			fmt.Println("Всего хорошего!")
			return
		}

		workGame()
		fmt.Println("Сыграть ещё раз? (y/n)")
		answer, _ := reader.ReadString('\n')
		if strings.TrimSpace(answer) != "y" {
			break
		}

	}
}

func workGame() {
	fmt.Println("Начнем игру")
	fmt.Println()

	//------------------
	word, err := repository.GetRandomWord("input.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Случайное слово:", word)
	fmt.Println(word)
	//------------------

	array := []rune(word)

	fmt.Println(view.FindTableWithLetters(array))

	//---создание пустого массива------

	secretRunes := []rune(word)
	arrayForGuessing := make([]rune, len(secretRunes))
	for y := range arrayForGuessing {
		arrayForGuessing[y] = '_'
	}

	guessedLetters := map[rune]bool{} // карта дубликатов букв

	//---запуск цыкла для запрса букв----

	countErrors := 0

	for i := 0; !service.IsWordGuessed(arrayForGuessing) && countErrors < limitErrors; i++ {

		fmt.Println(view.FindTableWithLetters(arrayForGuessing))
		fmt.Println("Введите букву: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len([]rune(input)) != 1 {
			fmt.Println("Пожалуйста, введите только одну букву.")
			continue
		}

		letter := []rune(input)[0]
		found := false

		if guessedLetters[letter] {
			fmt.Println("Вы уже вводили эту букву.")
			continue
		}

		guessedLetters[letter] = true

		for i, r := range secretRunes {
			if r == letter {
				arrayForGuessing[i] = r
				found = true
			}
		}

		if !found {
			view.FindImageHangman(countErrors)
			countErrors++
			fmt.Println("Такой буквы нет.")

		}

	}

	if service.IsWordGuessed(arrayForGuessing) {
		fmt.Println("Поздравляем! Вы отгадали слово:")
		fmt.Println(view.FindTableWithLetters(arrayForGuessing))
	} else {
		fmt.Println("Вы проиграли. Загаданное слово было:", word)
	}
}
