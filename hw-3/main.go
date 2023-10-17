package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/kljensen/snowball"
)

type Word struct {
	Name   string
	Repeat uint
}

// Возвращет топ-10 встречающихся слов
func BestRepeats(s string) []Word {

	wordsMap := map[string]uint{}

	//Убираем все лишние спец символы
	regex := regexp.MustCompile("[.,;:!?]")
	stringWithoutSimbols := regex.ReplaceAllString(s, "")

	//Разбиваем строку на слова
	wordsSplit := strings.Split(stringWithoutSimbols, " ")

	//Проходим по каждому слову и добавляем совадение
	for _, r := range wordsSplit {
		stemmed, err := snowball.Stem(r, "russian", true)
		if err == nil {
			wordsMap[stemmed]++
		}
	}

	var words []Word

	//Мы проходим по всей мапе и заполняем массив данными о кол-ве повторений
	for word, repeatCount := range wordsMap {
		words = append(words, Word{Name: word, Repeat: repeatCount})
	}

	//Сортируем массив структур по параметру Repeat
	sort.Slice(words, func(i, j int) bool {
		return words[i].Repeat > words[j].Repeat
	})
	if len(words) < 5 {
		return words
	}
	return words[:5]
}

func main() {
	fmt.Println(BestRepeats("Ты очень красивый, я немного влюбилась. Как ты считаешь, я красивая? Если что, я свободна сегодня, можем встретиться, если ты не занят."))
}
