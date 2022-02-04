package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const backslash rune = 92

var (
	ErrInvalidStringFirstElement   = errors.New("invalid string: first element can't be a digit")
	ErrInvalidStringDigitsSequence = errors.New("invalid string: only correctly escaped digits in a raw are valid")
	ErrInvalidStringEscapeSequence = errors.New("invalid string: only backspace or digit can be escaped")
)

func Unpack(data string) (string, error) {
	dataRunes := []rune(data)

	dataRunesLen := len(dataRunes)
	if dataRunesLen == 0 {
		return "", nil
	}

	dataIsValid, err := isDataValid(dataRunes, dataRunesLen)
	if !dataIsValid {
		return "", err
	}

	sb := strings.Builder{}
	for i := 0; i < dataRunesLen; i++ {
		if i == dataRunesLen-1 {
			if !unicode.IsDigit(dataRunes[i]) {
				sb.WriteRune(dataRunes[i])
			}
			break
		}

		if dataRunes[i] == backslash {
			processNextBackslash(dataRunes, i, &sb, &i)
			continue
		}

		if unicode.IsDigit(dataRunes[i+1]) {
			processNextDigit(dataRunes[i], dataRunes[i+1], &sb, &i)
			continue
		}

		sb.WriteRune(dataRunes[i])
	}

	return sb.String(), nil
}

func processNextDigit(currentRune rune, nextRune rune, sb *strings.Builder, runeIndex *int) {
	sb.WriteString(repeatRune(currentRune, nextRune))
	*runeIndex++
}

func processNextBackslash(dataRunes []rune, currentIndex int, sb *strings.Builder, runeIndex *int) {
	dataRunesLen := len(dataRunes)

	if currentIndex+2 < dataRunesLen {
		if unicode.IsDigit(dataRunes[currentIndex+2]) {
			repeatedRune := repeatRune(dataRunes[currentIndex+1], dataRunes[currentIndex+2])
			sb.WriteString(repeatedRune)
			*runeIndex += 2
		} else {
			sb.WriteRune(dataRunes[currentIndex+1])
			*runeIndex++
		}

		return
	}

	if currentIndex+2 == dataRunesLen {
		sb.WriteRune(dataRunes[currentIndex+1])
		return
	}
}

func repeatRune(currentRune rune, nextRune rune) string {
	runeRepeatCount, _ := strconv.Atoi(string(nextRune))
	result := ""
	if runeRepeatCount > 0 {
		result = strings.Repeat(string(currentRune), runeRepeatCount)
	}
	return result
}

func isDataValid(dataRunes []rune, dataRunesLen int) (bool, error) {
	if unicode.IsDigit(dataRunes[0]) {
		return false, ErrInvalidStringFirstElement
	}

	if dataRunes[dataRunesLen-1] == backslash {
		return false, ErrInvalidStringEscapeSequence
	}

	for i := 0; i < dataRunesLen; i++ {
		if dataRunes[i] == backslash && dataRunes[i+1] != backslash && !unicode.IsDigit(dataRunes[i+1]) {
			return false, ErrInvalidStringEscapeSequence
		}

		if i > 0 {
			if i+1 < dataRunesLen &&
				unicode.IsDigit(dataRunes[i]) && unicode.IsDigit(dataRunes[i+1]) && dataRunes[i-1] != backslash {
				return false, ErrInvalidStringDigitsSequence
			}

			if i+2 < dataRunesLen &&
				dataRunes[i] == backslash && dataRunes[i-1] == backslash &&
				unicode.IsDigit(dataRunes[i+1]) && unicode.IsDigit(dataRunes[i+2]) {
				return false, ErrInvalidStringDigitsSequence
			}
		}
	}

	return true, nil
}
