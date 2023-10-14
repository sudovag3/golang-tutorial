package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func UnpackString(s string) string {
	var last_rune rune
	var last_number int
	var lengthString int = utf8.RuneCountInString(s)
	var resultString strings.Builder
	var iter int = 1
	for _, symbol := range s {
		if unicode.IsDigit(symbol) {
			if last_rune == rune(0) {
				io.WriteString(os.Stderr,
					"Invalid String")
				return ""
			} else {
				number, err := strconv.Atoi(string(symbol))
				if err != nil {
					log.Fatal(err)
				}
				for i := 0; i < number; i++ {
					resultString.WriteRune(last_rune)
				}
				// int(symbol) раз записать руну
				last_rune = rune(0)
				last_number = number
			}
		} else {
			if iter == lengthString {
				if last_number == 0 {
					resultString.WriteRune(last_rune)
				}
				resultString.WriteRune(symbol)
			} else if last_rune == rune(0) {
				last_rune = symbol
			} else {
				resultString.WriteRune(last_rune)
				last_rune = symbol
			}
		}
		iter++
	}
	return resultString.String()
}

func main() {
}
