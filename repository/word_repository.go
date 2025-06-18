package repository

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
)

// GetRandomWord читает файл и возвращает случайное слово
func GetRandomWord(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineWords := strings.Fields(line)
		words = append(words, lineWords...)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "", errors.New("no words found in file")
	}

	// инициализация генератора случайных чисел

	//rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))

	return words[randomIndex], nil
}
