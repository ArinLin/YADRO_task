package main

import (
	"bufio"
	"fmt"
	"github.com/kljensen/snowball"
	"os"
	"strings"
)

func main() {
	// Открываю список стоп-слов
	file, err := os.Open("stop_words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Читаю стоп-слова из файла и добавляю их в map
	stopWordsMap := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWordsMap[strings.ToLower(scanner.Text())] = struct{}{}
	}

	// Читаю ввод пользователя из консоли
	fmt.Println("Enter text to normalize:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Перебираю слова введенные пользователем, если слово не содержится в стоп-листе, то нормализую и вывожу
	for _, word := range strings.Fields(input) {
		toLowerCase := strings.ToLower(word)
		_, isStopWods := stopWordsMap[toLowerCase]
		if !isStopWods {
			stemmed, err := snowball.Stem(toLowerCase, "english", true)
			if err == nil {
				fmt.Print(stemmed + " ")
			}
		}
	}
	fmt.Println()
}
