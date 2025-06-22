package game

import (
	"math/rand"
)

var wordList = [...]string{
	"gopher", "program", "developer", "function", "variable", "interface", "channel",
}

func FindWord() string {
	return wordList[rand.Intn(len(wordList))]
}
