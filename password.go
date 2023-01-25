package password

import (
	"errors"
	"strings"
)

var (
	ErrLength8                = errors.New("password should be at least 8 characters long")
	ErrMissingNumeric         = errors.New("password should at least have 2 numeric characters")
	ErrMissingAlphabetic      = errors.New("password should at least have 4 alphabetic characters")
	ErrMissingSpecial         = errors.New("password should at least have 1 special character")
	ErrMissingUppercase       = errors.New("password should at least have 1 uppercase character")
	ErrEqualAdjacent          = errors.New("password should not contain adjacent characters with the same value")
	ErrInvalidCharCombination = errors.New("password contains an invalid combination of characters: 'asdf', 'qwerty', '1234' or '98765'")
	ErrConsecutive            = errors.New("password contains values consecutive to each other, 1234, 3456, abcd, efgh")
)

// Validate returns an error if the given password is not valid according to a set of rules defined by this package.
//
//	Validation rules:
//		- password should be at least 8 characters long
//		- password should at least have 2 numeric characters
//		- password should at least have 4 alphabetic characters
//		- password should at least have 1 special character
//		- password should at least have 1 uppercase character
//		- password should not contain adjacent characters with the same value
//		- password contains an invalid combination of characters: 'asdf', 'qwerty', '1234' or '98765'
//		- password contains values consecutive to each other, 1234, 3456, abcd, efgh
func Validate(password string) error {
	if len(password) < 8 {
		return ErrLength8
	}

	var numerics int
	var alphabetic int
	var specials int
	var uppercase int
	var buffer [4]rune
	var equal int
	var consecutive int

	for i, p := range password {
		buffer[i%4] = p

		if isNumericChar(p) {
			numerics++
		}
		if hasUppercase := isAlphabeticUppercase(p); hasUppercase || isAlphabeticLowercase(p) {
			if hasUppercase {
				uppercase++
			}
			alphabetic++
		}
		if isSpecialChar(p) {
			specials++
		}

		for _, c := range buffer {
			if isSpecialChar(p) {
				continue
			}
			if p == c {
				equal++
			}
			if p-c <= 3 && p-c >= -3 {
				consecutive++
			}
		}

		if equal == 4 {
			return ErrEqualAdjacent
		}
		if consecutive == 4 {
			return ErrConsecutive
		}
		equal = 0
		consecutive = 0
	}
	if numerics < 2 {
		return ErrMissingNumeric
	}
	if alphabetic < 4 {
		return ErrMissingAlphabetic
	}
	if uppercase < 1 {
		return ErrMissingUppercase
	}
	if specials < 1 {
		return ErrMissingSpecial
	}

	if strings.Index(password, "asdf") != -1 {
		return ErrInvalidCharCombination
	}

	if strings.Index(password, "qwerty") != -1 {
		return ErrInvalidCharCombination
	}

	return nil
}

// isAlphabeticUppercase returns true if the given char is a lowercase character.
func isAlphabeticLowercase(c rune) bool {
	return c >= 'a' && c <= 'z'
}

// isAlphabeticUppercase returns true if the given char is an uppercase character.
func isAlphabeticUppercase(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

// isNumericChar returns true when the given char is a number.
func isNumericChar(c rune) bool {
	return c >= '0' && c <= '9'
}

// isSpecialChar determines if the given char is a valid special character according to the ASCII table.
// Some chars have been omitted in order to keep this function simple.
func isSpecialChar(c rune) bool {
	return (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~')
}
