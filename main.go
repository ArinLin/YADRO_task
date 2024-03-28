package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kljensen/snowball"
)

const (
	rus = "russian"
	eng = "english"
)

func main() {
	var sentence string
	var language string
	flag.StringVar(&sentence, "s", "", "sentence for stemming")
	flag.StringVar(&language, "l", "", "stemming language")
	flag.Parse()
	if sentence == "" {
		fmt.Println("Enter sentence for stemming")
		return
	}
	if language == "" {
		language = "english"
	}

	stopWordsMap, err := createStopWordsMap(language)
	if err != nil {
		fmt.Printf("error create stop words map: %s/n", err)
		return
	}
	normalizedSentence, err := normalizeSentence(sentence, language, stopWordsMap)
	if err != nil {
		fmt.Printf("error normalize sentence: %s/n", err)
		return
	}
	fmt.Println(normalizedSentence)
}

// Функция для создания мапы стоп-слов
func createStopWordsMap(language string) (map[string]struct{}, error) {
	var fileName string
	switch language {
	case rus:
		fileName = "stop_words_rus.txt"
	case eng:
		fileName = "stop_words_eng.txt"
	default:
		return nil, errors.New("unavaliable language")
	}
	// Открываю список стоп-слов
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Читаю стоп-слова из файла и добавляю их в map
	stopWordsMap := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWordsMap[strings.ToLower(scanner.Text())] = struct{}{}
	}
	return stopWordsMap, nil
}

// Функция для нормализации предложения
func normalizeSentence(sentence, language string, stopWordsMap map[string]struct{}) (string, error) {
	var normalizedWords []string
	// Перебираю слова введенные пользователем, если слово не содержится в стоп-листе, то нормализую и вывожу
	for _, word := range strings.Fields(sentence) {
		toLowerCase := strings.ToLower(word)
		if _, isStopWord := stopWordsMap[toLowerCase]; !isStopWord {
			stemmed, err := snowball.Stem(toLowerCase, language, true)
			if err != nil {
				return "", err
			}
			normalizedWords = append(normalizedWords, stemmed)
		}
	}
	return strings.Join(normalizedWords, " "), nil
}
