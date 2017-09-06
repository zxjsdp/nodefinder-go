package nodefindergo

import (
	"fmt"
	"log"
	"strings"

	"github.com/zxjsdp/nodefinder-go/utils"
)

const (
	// Normal calibration info or clade label type
	CALI_OR_CLADE_LABEL_TYPE = "cali_or_clade_label"
	// Branch label type
	BRANCH_LABEL_TYPE = "branchLabel"

	// Blank name for branch label, which needs only one species name
	BLANK_NAME_B = ""

	DEFAULT_CALI_OR_CLADE_LABEL_DESCRIPTION = "Normal calibration or clade label."
	DEFAULT_BRANCH_LABEL_DESCRIPTION        = "Branch label description"

	USAGE = "\n\n[ USAGE ] nodefindergo -input input.nwk " +
		"-config config.txt -output output.nwk\n[ USAGE ] Please try `nodefinder-go -h` for help.\n"
)

var (
	NOT_TREE_NAME_SYMBOLS []rune = []rune{',', ';', ')', '"', '#', '$', '@', '>', '<'}
	print                        = fmt.Println
)

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

func GetInsertionList(cleanTreeStr, name string) []int {
	insertionList := []int{}
	currentIndex := strings.Index(cleanTreeStr, name)
	stack := []rune{}
	strLen := len(cleanTreeStr)

	for currentIndex < strLen {
		if cleanTreeStr[currentIndex] == '(' {
			stack = append(stack, '(')
		} else if cleanTreeStr[currentIndex] == ')' {
			if len(stack) == 0 {
				insertionList = append(insertionList, currentIndex+1)
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
	insertionListA := GetInsertionList(cleanTreeStr, nameA)
	insertionListB := GetInsertionList(cleanTreeStr, nameB)

	utils.Reverse(insertionListA)
	utils.Reverse(insertionListB)

	longerInsertionList, shorterInsertionList := utils.FindLongerAndShorterArray(
		insertionListA, insertionListB)

	for index, insertionPoint := range shorterInsertionList {
		if index == len(shorterInsertionList)-1 {
			indexOfTMRCA = insertionPoint
		}
		if shorterInsertionList[index] != longerInsertionList[index] {
			indexOfTMRCA = shorterInsertionList[index-1]
			break
		}
	}
	return indexOfTMRCA
}

func SingleCalibration(rawTreeStr, nameA, nameB, caliInfo string) string {
	cleanTreeStr := GetCleanTreeStr(rawTreeStr)
	insertionPoint := GetIndexOfTMRCA(cleanTreeStr, nameA, nameB)

	leftPart := cleanTreeStr[:insertionPoint]
	rightPart := cleanTreeStr[insertionPoint:]
	cleanTreeStr = leftPart + caliInfo + rightPart

	return cleanTreeStr
}

func AddSingleBranchLabel(rawTreeStr, nameA, branchLabel string) string {
	cleanTreeStr := GetCleanTreeStr(rawTreeStr)
	insertionPoint := GetRightIndexOfName(cleanTreeStr, nameA)

	leftPart := cleanTreeStr[:insertionPoint]
	rightPart := cleanTreeStr[insertionPoint:]
	cleanTreeStr = leftPart + branchLabel + rightPart

	return cleanTreeStr
}

func MultipleCalibration(rawTreeStr string, calibrations []Calibration) string {
	for index, cali := range calibrations {
		if cali.CaliType == CALI_OR_CLADE_LABEL_TYPE {
			rawTreeStr = SingleCalibration(rawTreeStr, cali.NameA, cali.NameB, cali.CaliInfo)
		} else if cali.CaliType == BRANCH_LABEL_TYPE {
			rawTreeStr = AddSingleBranchLabel(rawTreeStr, cali.NameA, cali.CaliInfo)
		} else {
			log.Panic("Invalid calibration!")
		}
		print(fmt.Sprintf("%3d: %s", index, rawTreeStr))
	}
	return strings.Replace(rawTreeStr, ",", ", ", -1)
}

func ParseConfig(configFileName string) []Calibration {
	calibrations := []Calibration{}
	lines := utils.ReadLines(configFileName)
	for index, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			elements := strings.Split(line, ",")
			if len(elements) == 3 {
				elements = utils.CleanElements(elements)
				currentCalibration := Calibration{index, CALI_OR_CLADE_LABEL_TYPE,
					elements[0], elements[1], elements[2],
					DEFAULT_CALI_OR_CLADE_LABEL_DESCRIPTION}
				calibrations = append(calibrations, currentCalibration)
			} else if len(elements) == 2 {
				elements = utils.CleanElements(elements)
				currentCalibration := Calibration{index, BRANCH_LABEL_TYPE,
					elements[0], BLANK_NAME_B, elements[1],
					DEFAULT_BRANCH_LABEL_DESCRIPTION}
				calibrations = append(calibrations, currentCalibration)
			} else {
				log.Fatal(fmt.Sprintf("Invalid config line: %d", (index + 1)))
			}
		}
	}
	return calibrations
}

func Test() {
	//var calibrations []Calibration
	rawTreeStr := "((a,((b,c),(ddd,e))),(f,g));"
	//
	//calibrations = []Calibration{
	//	Calibration{0, CALI_OR_CLADE_LABEL_TYPE, "b", "ddd", ">0.1<0.2", "First calibration"},
	//	Calibration{1, CALI_OR_CLADE_LABEL_TYPE, "b", "c", ">0.03<0.05", "Second calibration"},
	//	Calibration{2, BRANCH_LABEL_TYPE, "f", "", "#3", "First branchLabel"},
	//}

	//print(MultipleCalibration(rawTreeStr, calibrations))
	calibrations := ParseConfig(
		"/Users/jin/go/src/github.com/zxjsdp/nodefinder-go/nodefindergo_tests/calibration.txt")
	print(MultipleCalibration(rawTreeStr, calibrations))
}
