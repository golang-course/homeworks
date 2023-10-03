package hw3

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*func StringUnpacking(text string) (string, error) {
	var newstring strings.Builder
	newstring.Grow(3 * len(text))
	var prevvalue string
	//slice := []rune(s)
	for i, value := range text {
		currentvalue := string(value)
		intvalue, err := strconv.Atoi(currentvalue)
		if err != nil {
			intvalue = 1
			prevvalue = currentvalue
		}
		if i == 0 && err == nil {
			return "", errors.New("starts with a number!")
		}
		if err == nil {
			intvalue -= 1
		}
		newstring.WriteString(strings.Repeat(prevvalue, intvalue))
	}
	return fmt.Sprint(newstring), nil
}*/

const escapeRune = '\\'

func StringUnpacking(line string) (string, error) {
	var result strings.Builder
	var counter int
	var currentRune rune
	runeSlice := []rune(line)
	runeLen := len(runeSlice)

	for i := 0; i < runeLen; i++ {
		currentRune = runeSlice[i]
		counter = 1

		if unicode.IsDigit(currentRune) && i == 0 {
			return "", errors.New("error - starts with a number")

		}
		if currentRune == escapeRune {
			if i+1 > runeLen {
				return "", errors.New("error - end of len")
			}
			currentRune = runeSlice[i+1]
			if !unicode.IsDigit(currentRune) && currentRune != escapeRune {
				return "", errors.New("error - after \\ must be a number or a \\")
			}
			i++
		}
		if i+1 < runeLen && unicode.IsDigit(runeSlice[i+1]) {
			if i+2 < runeLen {
				if unicode.IsDigit(runeSlice[i+2]) {
					return "", errors.New("error - only numbers are allowed")
				}
			}
			counter, _ = strconv.Atoi(string(runeSlice[i+1]))
			i++
		}
		result.WriteString(strings.Repeat(string(currentRune), counter))
	}
	return result.String(), nil
}
