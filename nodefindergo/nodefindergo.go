package nodefindergo

import (
	"github.com/zxjsdp/nodefinder-go/utils"
	"strings"
)

var NOT_TREE_NAME_SYMBOLS []rune = []rune {',', ';', ')', '"', '#', '$', '@', '>', '<'}

func GetCleanTreeStr(rawTreeStr string) string {
	newTreeStr := utils.RemoveChar(rawTreeStr, ' ')
	newTreeStr = utils.RemoveChar(newTreeStr, '\n')
	newTreeStr = utils.RemoveChar(newTreeStr, '\t')

	return newTreeStr
}

func GetRightIndexOfName(cleanTreeStr, name string) int {
	leftIndexOfName := strings.Index(cleanTreeStr, name)
	for !utils.CheckRuneInRunesV2(NOT_TREE_NAME_SYMBOLS, []rune(cleanTreeStr)[leftIndexOfName]) {
		leftIndexOfName += 1
	}
	return leftIndexOfName
}

func getInsertionList(cleanTreeStr, name string) []int {
	insertionList := []int{}
	currentIndex := strings.Index(cleanTreeStr, name)
	stack := []rune{}
	strLen := len(cleanTreeStr)

	for currentIndex < strLen {
		if cleanTreeStr[currentIndex] == '(' {
			stack = append(stack, '(')
		} else if cleanTreeStr[currentIndex] == ')' {
			if len(stack) == 0 {
				insertionList = append(insertionList, currentIndex + 1)
			} else {
				stack = append(stack[:len(stack)-1])
			}
		}
		currentIndex += 1
	}

	return insertionList
}

func GetIndexOfTMRCA(cleanTreeStr, nameA, nameB string) int {
	var indexOfTMRCA int
	insertionListA := getInsertionList(cleanTreeStr, nameA)
	insertionListB := getInsertionList(cleanTreeStr, nameB)

	utils.Reverse(insertionListA)
	utils.Reverse(insertionListB)

	longerInsertionList, shorterInsertionList := utils.FindLongerAndShorterArray(
		insertionListA, insertionListB)

	for index, insertionPoint := range shorterInsertionList {
		if index == len(shorterInsertionList) - 1 {
			indexOfTMRCA = insertionPoint
		}
		if shorterInsertionList[index] != longerInsertionList[index] {
			indexOfTMRCA = shorterInsertionList[index - 1]
			break
		}
	}
	return indexOfTMRCA
}