package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var inputString string

	fmt.Print("Введите строку, в которой хотите поменять слова местами: ")

	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	inputString = strings.TrimPrefix(inputString, "\n")

	var resultString string
	
	resultString = strings.ReplaceAll(SwapWords(inputString), "\n", "")
	fmt.Printf("Результат: %s", resultString)
}

func SwapWords(inputString string) string {
	// Меняет слова местами в строке.	

	runeArr := []rune(inputString)

	for i := 0; i <= (len(runeArr) - 1) / 2; i++ {
		temp := runeArr[len(runeArr) - 1 - i]
		runeArr[len(runeArr) - 1 - i] = runeArr[i]
		runeArr[i] = temp
	}

	first := 0
	last := 0
	spaceCount := 0
	
	// привет тебе
	// ебет тевирп first = 0 last = 4 - 1 wordLength = 4 - 1 - 0 + 1 = 4
	// тебе првиет

	for i := 0; i <= len(runeArr) - 1; i++ {
		if string(runeArr[i]) == " " || i + 1 > len(runeArr) - 1{
		
			if i + 1 > len(runeArr) - 1 {
				last = i // индекс последней буквы в слове
			} else {
				last = i - 1
			}
			
			wordLength := last - first + 1	
			for j := 0; j < wordLength / 2; j++ { // цикл по слову
				temp := runeArr[last - j]
				runeArr[last - j] = runeArr[first + j]
				runeArr[first + j] = temp
			}
			first = i + 1
			spaceCount++
		}
	}

	if spaceCount == 0 { // если нет пробелов, просто переворачиваем слово.
		for i := 0; i <= (len(runeArr) - 1) / 2; i++ {
			temp := runeArr[len(runeArr) - 1 - i]
			runeArr[len(runeArr) - 1 - i] = runeArr[i]
			runeArr[i] = temp
		}
	}

	return string(runeArr) // Пока просто затычка
}
