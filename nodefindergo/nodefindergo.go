package nodefindergo

import (
	"github.com/zxjsdp/nodefinder-go/utils"
	"strings"
	"fmt"
	"log"
)

var (
	NOT_TREE_NAME_SYMBOLS []rune = []rune {',', ';', ')', '"', '#', '$', '@', '>', '<'}
	print = fmt.Println
)

func getCleanTreeStr(rawTreeStr string) string {
	newTreeStr := utils.RemoveChar(rawTreeStr, ' ')
	newTreeStr = utils.RemoveChar(newTreeStr, '\n')
	newTreeStr = utils.RemoveChar(newTreeStr, '\t')

	return newTreeStr
}

func getRightIndexOfName(cleanTreeStr, name string) int {
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

func getIndexOfTMRCA(cleanTreeStr, nameA, nameB string) int {
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

func singleCalibration(rawTreeStr, nameA, nameB, caliInfo string) string {
	cleanTreeStr := getCleanTreeStr(rawTreeStr)
	insertionPoint := getIndexOfTMRCA(cleanTreeStr, nameA, nameB)

	leftPart := cleanTreeStr[:insertionPoint]
	rightPart := cleanTreeStr[insertionPoint:]
	cleanTreeStr = leftPart + caliInfo + rightPart

	return cleanTreeStr
}

func addSingleBranchLabel(rawTreeStr, nameA, branchLabel string) string {
	cleanTreeStr := getCleanTreeStr(rawTreeStr)
	insertionPoint := getRightIndexOfName(cleanTreeStr, nameA)

	leftPart := cleanTreeStr[:insertionPoint]
	rightPart := cleanTreeStr[insertionPoint:]
	cleanTreeStr = leftPart + branchLabel + rightPart

	return cleanTreeStr
}

func MultipleCalibration(rawTreeStr string, calibrations []Calibration) string {
	for index, cali := range calibrations {
		if cali.caliType == "calibration" || cali.caliType == "cladeLabel" {
			rawTreeStr = singleCalibration(rawTreeStr, cali.nameA, cali.nameB, cali.caliInfo)
		} else if cali.caliType == "branchLabel" {
			rawTreeStr = addSingleBranchLabel(rawTreeStr, cali.nameA, cali.caliInfo)
		} else {
			log.Panic("Invalid calibration!")
		}
		print(fmt.Sprintf("%3d: %s", index, rawTreeStr))
	}
	return strings.Replace(rawTreeStr, ",", ", ", -1)
}

func Test() {
	var calibrations []Calibration
	rawTreeStr := "((a,((b,c),(ddd,e))),(f,g));"

	calibrations = []Calibration{
		Calibration{0, "calibration", "b", "ddd", ">0.1<0.2", "First calibration"},
		Calibration{1, "cladeLabel", "b", "c", ">0.03<0.05", "Second calibration"},
		Calibration{2, "branchLabel", "f", "", "#3", "First branchLabel"},
	}

	print(MultipleCalibration(rawTreeStr, calibrations))
}
