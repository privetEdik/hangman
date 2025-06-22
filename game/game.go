package game

import "slices"

type Hangman struct {
	Word         string
	Guessed      []bool
	AttemptsLeft int
	TriedLetters map[rune]bool
}

// создвть сущность самой игры . передать слово и  количество попыток

func NewGame(word string, attempts int) *Hangman {
	return &Hangman{
		Word:         word,                    // само  слово
		Guessed:      make([]bool, len(word)), //сделать срез,масив. буду отмечать какие позицыи  угаданы
		AttemptsLeft: attempts,                // попытки
		TriedLetters: make(map[rune]bool)}     // карта для учета уже названых букв
}

func (h *Hangman) GuessLetter(letter rune) string {
	if h.TriedLetters[letter] {
		return "You already tried that letter"
	}
	h.TriedLetters[letter] = true
	correct := false
	for i, ch := range h.Word {
		if ch == letter {
			h.Guessed[i] = true
			correct = true
		}
	}
	if !correct {
		h.AttemptsLeft--
		return "Incorrect guess."
	}
	return "Correct!"
}

func (h *Hangman) IsOver() bool {
	return h.AttemptsLeft == 0 || h.IsWon()
}

/*func (h *Hangman) IsWon() bool {
	for _, ok := range h.Guessed {
		if !ok {
			return false
		}
	}
	return true
}*/

func (h *Hangman) IsWon() bool {
	return !slices.Contains(h.Guessed, false)
}
