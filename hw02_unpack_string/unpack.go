package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const backslash rune = 92

var (
	ErrInvalidStringDigitFirst          = errors.New("invalid string – first symbol can't be a digit")
	ErrInvalidStringDigitInRawCount     = errors.New("invalid string – can't have two or more digits in a row")
	ErrInvalidStringBackslashSequence   = errors.New("invalid string – after backslash can stand only digit or \\")
	ErrInvalidStringBackslashInRawCount = errors.New("invalid string – max two backslashes in a raw are valid")
)

func Unpack(data string) (string, error) {
	dataRunes := []rune(data)

	dataRunesLen := len(dataRunes)
	if dataRunesLen == 0 {
		return "", nil
	}

	dataIsValid, err := isDataValid(dataRunes)
	if !dataIsValid {
		return "", err
	}

	builder := strings.Builder{}
	isPreviousStepRepeated := false

	for i := 0; i < dataRunesLen; i++ {
		if i == 0 {
			processRune(&builder, 0, dataRunes[i], dataRunes[i+1], &isPreviousStepRepeated)
			continue
		}

		if i == dataRunesLen-1 {
			processRune(&builder, dataRunes[i-1], dataRunes[i], 0, &isPreviousStepRepeated)
			break
		}

		processRune(&builder, dataRunes[i-1], dataRunes[i], dataRunes[i+1], &isPreviousStepRepeated)
	}

	return builder.String(), nil
}

func processRune(b *strings.Builder, previousRune rune, currentRune rune, nextRune rune, isPreviousStepRepeated *bool) {
	if (currentRune == backslash || unicode.IsDigit(currentRune)) && previousRune != backslash {
		*isPreviousStepRepeated = false

		return
	}

	if currentRune == backslash && previousRune == backslash && !unicode.IsDigit(nextRune) && !*isPreviousStepRepeated {
		b.WriteRune(currentRune)
		*isPreviousStepRepeated = false

		return
	}

	if unicode.IsDigit(currentRune) && previousRune == backslash &&
		!unicode.IsDigit(nextRune) && !*isPreviousStepRepeated {
		b.WriteRune(currentRune)
		*isPreviousStepRepeated = false

		return
	}

	if !unicode.IsDigit(nextRune) && !unicode.IsDigit(currentRune) {
		b.WriteRune(currentRune)
		*isPreviousStepRepeated = false

		return
	}

	if unicode.IsDigit(nextRune) {
		runeRepeatCount, _ := strconv.Atoi(string(nextRune))
		if runeRepeatCount > 0 {
			repeatedString := strings.Repeat(string(currentRune), runeRepeatCount)
			b.WriteString(repeatedString)
		}
		*isPreviousStepRepeated = true

		return
	}
}

func isDataValid(dataRunes []rune) (bool, error) {
	dataIsValid, err := isFirstRuneValid(dataRunes[0])
	if !dataIsValid {
		return false, err
	}

	dataIsValid, err = isRuneAfterBackslashValid(dataRunes)
	if !dataIsValid {
		return false, err
	}

	dataIsValid, err = isDigitSequenceValid(dataRunes)
	if !dataIsValid {
		return false, err
	}

	dataIsValid, err = isBackslashSequenceValid(dataRunes)
	if !dataIsValid {
		return false, err
	}

	return true, nil
}

func isFirstRuneValid(firstRune rune) (bool, error) {
	if unicode.IsDigit(firstRune) {
		return false, ErrInvalidStringDigitFirst
	}

	return true, nil
}

func isRuneAfterBackslashValid(dataRunes []rune) (bool, error) {
	for i := 0; i < len(dataRunes)-1; i++ {
		if dataRunes[i] == backslash && !unicode.IsDigit(dataRunes[i+1]) && dataRunes[i+1] != backslash {
			return false, ErrInvalidStringBackslashSequence
		}
	}
	if dataRunes[len(dataRunes)-1] == backslash {
		return false, ErrInvalidStringBackslashSequence
	}

	return true, nil
}

func isDigitSequenceValid(dataRunes []rune) (bool, error) {
	for i := 1; i < len(dataRunes)-1; i++ {
		if dataRunes[i-1] != backslash && unicode.IsDigit(dataRunes[i]) && unicode.IsDigit(dataRunes[i+1]) {
			return false, ErrInvalidStringDigitInRawCount
		}

		if i-2 >= 0 {
			if dataRunes[i-2] == backslash && dataRunes[i-1] == backslash &&
				unicode.IsDigit(dataRunes[i]) && unicode.IsDigit(dataRunes[i+1]) {
				return false, ErrInvalidStringDigitInRawCount
			}
		}
	}

	return true, nil
}

func isBackslashSequenceValid(dataRunes []rune) (bool, error) {
	for i := 1; i < len(dataRunes)-2; i++ {
		if dataRunes[i-1] == backslash && dataRunes[i] == backslash && dataRunes[i+1] == backslash {
			return false, ErrInvalidStringBackslashInRawCount
		}
	}

	return true, nil
}
