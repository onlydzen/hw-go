package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "\t3", expected: "\t\t\t"},
		{input: "日本3語", expected: "日本本本語"},
		{input: "t0", expected: ""},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `\3`, expected: `3`},
		{input: `\\3`, expected: `\\\`},
		{input: `\\0`, expected: ``},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidStringFirstDigit(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "0"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidStringDigitFirst), "actual error %q", err)
		})
	}
}

func TestUnpackInvalidStringDigitsInRaw(t *testing.T) {
	invalidStrings := []string{"we45", "aaa\n10b", "asd日213", `\\31`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidStringDigitInRawCount), "actual error %q", err)
		})
	}
}

func TestUnpackInvalidStringBackslashes(t *testing.T) {
	invalidStrings := []string{`qw\ne`, `\qwne`, `qwne\`, `qwe\4\5\`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidStringBackslashSequence), "actual error %q", err)
		})
	}
}

func TestUnpackInvalidStringBackslashesInRowCount(t *testing.T) {
	invalidStrings := []string{`qwe\\\3`, `\\\\8wne`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidStringBackslashInRawCount), "actual error %q", err)
		})
	}
}
