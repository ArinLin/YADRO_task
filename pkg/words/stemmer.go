package words

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"github.com/kljensen/snowball"
)

const fileName = "stop_words_eng.txt"

// Функция для создания мапы стоп-слов
func CreateStopWordsMap() (map[string]struct{}, error) {
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
func NormalizeSentence(sentence string, stopWordsMap map[string]struct{}) ([]string, error) {
	cleanSentence := cleanString(sentence)
	var normalizedWords []string
	// Перебираю слова введенные пользователем, если слово не содержится в стоп-листе, то нормализую и вывожу
	for _, word := range strings.Fields(cleanSentence) {
		toLowerCase := strings.ToLower(word)
		if _, isStopWord := stopWordsMap[toLowerCase]; !isStopWord {
			stemmed, err := snowball.Stem(toLowerCase, "english", true)
			if err != nil {
				return nil, err
			}
			normalizedWords = append(normalizedWords, stemmed)
		}
	}
	return normalizedWords, nil
}

func cleanString(input string) string {
	var result strings.Builder

	// Перебор символов в строке
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsSpace(r) || unicode.IsNumber(r) {
			// Если символ буква, цифра или пробел, добавляем его в результат
			result.WriteRune(r)
		}
	}

	// Преобразуем результат в строку и убираем лишние пробелы
	return strings.Join(strings.Fields(result.String()), " ")
}
