package nodefindergo_test

import (
	"github.com/zxjsdp/nodefinder-go/utils"
	"testing"
	"io/ioutil"
	"os"
	"reflect"
)

var (
	rawTree string = "((a, ((b, c), (ddd,\t e))), (f, g));\n\n"
	cleanTree string = "((a,((b,c),(ddd,e))),(f,g));"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// stringUtil.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func Test_CleanElements(t *testing.T) {
	dirtyElements := []string{" \tsome\n", " \ndirty", "elements  \t\n"}
	expectedCleanElements := []string{"some", "dirty", "elements"}

	if len(utils.CleanElements(dirtyElements)) != len(expectedCleanElements) {
		t.Error("CleanElements (stringUtil.go): Should remove blanks, tab, and newlines in string elements")
	}
}

func Test_RemoveChar(t *testing.T) {
	expectedCleanString := "((a, ((b, c), (ddd, e))), (f, g));\n\n"
	result := utils.RemoveChar(rawTree, '\t')

	if (result != expectedCleanString) {
		t.Error("RemoveChar (stringUtil.go): Should remove specific rune from string")
	}
}

func Test_RemoveBlankChars(t *testing.T) {
	expectedReplacedString := "((a,_((b,_c),_(ddd,__e))),_(f,_g));__"

	result := utils.ReplaceBlankChars(rawTree)

	if (result != expectedReplacedString) {
		t.Error("ReplaceBlankChars (stringUtil.go): Should remove all blank runes")
	}
}

func Test_CheckRuneExistsInString(t *testing.T) {
	subString := "a,((b,c)"

	if (!utils.CheckSubStringExistsInString(cleanTree, subString)) {
		t.Error("CheckSubStringExistsInString (stringUtil.go): Check substring in string failed")
	}
}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// stringUtil.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func Test_IOUtils(t *testing.T) {
	expectedLineNum := 3

	lines := utils.ReadLines("calibration.txt")

	if (len(lines) != expectedLineNum) {
		t.Error("ReadLines (ioUtil.go): Line number in file not match")
	}
}

func Test_ReadContent(t *testing.T) {
	expectedContent := "b, ddd, >0.1<0.2\nb, c, $5\nf, #3"

	result := utils.ReadContent("calibration.txt")

	if (result != expectedContent) {
		t.Error("ReadContent (ioUtil.go): Read content failed")
	}
}

func Test_ReadCleanContent(t *testing.T) {
	expectedCleanContent := cleanTree

	result := utils.ReadCleanContent("input.nwk", []rune{' ', '\t', '\n'})

	if (result != expectedCleanContent) {
		t.Error("ReadCleanContent (ioUtil.go): ReadCleanContent failed")
	}
}

func Test_CheckFileExists(t *testing.T) {
	utils.CheckFileExists("input.nwk", "description", "usage")
}

func Test_WriteContent(t *testing.T) {
	contentToBeWritten := "the fox jumped over the lazy dog"
	tmpFile, _ := ioutil.TempFile(os.TempDir(), "testWrite")
	defer os.Remove(tmpFile.Name())

	utils.WriteContent(tmpFile.Name(), contentToBeWritten)

	utils.CheckFileExists(tmpFile.Name(), "description", "usage")
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// sliceUtil.go
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func Test_CheckRuneInRunesV2_CheckTrue(t *testing.T) {
	runes := []rune{'*', '#', '$', '%'}
	runeInRunes := '*'

	if !(utils.CheckRuneInRunesV2(runes, runeInRunes)) {
		t.Error("CheckRuneInRunesV2 (sliceUtil.go): Rune in runes check failed (Rune in runes)")
	}
}

func Test_CheckRuneInRunesV2_CheckFalse(t *testing.T) {
	runes := []rune{'*', '#', '$', '%'}
	runeNotInRunes := '&'

	if (utils.CheckRuneInRunesV2(runes, runeNotInRunes)) {
		t.Error("CheckRuneInRunesV2 (sliceUtil.go): Rune in runes check failed (Rune not in runes)")
	}
}

func Test_CheckRuneInRunesV1_CheckTrue(t *testing.T) {
	runes := []rune{'*', '#', '$', '%'}
	runeInRunes := '*'

	if !(utils.CheckRuneInRunesV1(runes, runeInRunes)) {
		t.Error("CheckRuneInRunesV1 (sliceUtil.go): Rune in runes check failed (Rune in runes)")
	}
}

func Test_CheckRuneInRunesV1_CheckFalse(t *testing.T) {
	runes := []rune{'*', '#', '$', '%'}
	runeNotInRunes := '&'

	if (utils.CheckRuneInRunesV1(runes, runeNotInRunes)) {
		t.Error("CheckRuneInRunesV1 (sliceUtil.go): Rune in runes check failed (Rune not in runes)")
	}
}

func Test_Reverse(t *testing.T) {
	originalSlice := []int{2, 5, 3, 7, 1}
	expectedSlice := []int{1, 7, 3, 5, 2}

	utils.Reverse(originalSlice)

	if (!reflect.DeepEqual(originalSlice, expectedSlice)) {
		t.Error("Reverse (sliceUtil.go): Reverse int slice failed")
	}
}

func Test_FindLongerAndShorterArray(t *testing.T) {
	shorterArray := []int{2, 5, 1, 9}
	longerArray := []int{8, 2, 0, 5, 1, 9, 10}

	newLongerArray, newShorterArray := utils.FindLongerAndShorterArray(shorterArray, longerArray)

	if (!reflect.DeepEqual(newLongerArray, longerArray) ||
			!reflect.DeepEqual(newShorterArray, shorterArray)) {
		t.Error("FindLongerAndShorterArray (sliceUtil.go): Finder longer array and shorter array failed")
	}
}
