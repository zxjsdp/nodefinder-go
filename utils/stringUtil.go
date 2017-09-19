package utils

import (
	"strings"
	"unicode"
)

// CleanElements clean whitespaces in strings.
func CleanElements(strList []string) []string {
	for index, each := range strList {
		each = strings.TrimSpace(each)
		strList[index] = each
	}
	return strList
}

// RemoveChar removes rune from string.
func RemoveChar(str string, runeToBeRemoved rune) string {
	return strings.Map(func(r rune) rune {
		if r != runeToBeRemoved {
			return r
		}
		return -1
	}, str)
}

// ReplaceBlankChars replaces blank chars in string with _.
func ReplaceBlankChars(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '_'
		}
		return r
	}, str)
}

// CheckSubStringExistsInString checks whether substring exists in string.
func CheckSubStringExistsInString(str, subString string) bool {
	index := strings.Index(str, subString)
	if index >= 0 {
		return true
	}
	return false
}
