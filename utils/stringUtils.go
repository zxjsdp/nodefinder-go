package utils

import (
	"strings"
	"unicode"
)

func CleanElements(strList []string) []string {
	for index, each := range strList {
		each = strings.TrimSpace(each)
		strList[index] = each
	}
	return strList
}

func RemoveChar(str string, runeToBeRemoved rune) string {
	return strings.Map(func(r rune) rune {
		if r != runeToBeRemoved {
			return r
		}
		return 0
	}, str)
}

func RemoveBlankChars(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return '_'
		}
		return r
	}, str)
}

func CheckRuneExistsInString(str, subString string) bool {
	index := strings.Index(str, subString)
	if index >= 0 {
		return true
	}
	return false
}
