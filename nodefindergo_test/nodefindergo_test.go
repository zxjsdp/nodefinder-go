package nodefindergo

import (
	"github.com/zxjsdp/nodefinder-go/nodefindergo"
	"testing"
	"fmt"
	"reflect"
	"os"
	"path"
)

var (
	rawTree string = "((a, ((b, c), (ddd,\t e))), (f, g));\n\n"
	cleanTree string = "((a,((b,c),(ddd,e))),(f,g));"
	nameA string = "ddd"
	nameB string = "b"
	caliInfo string = ">0.1<0.2"
	branchLabel string = "@0.3"
)

func Test_getCleanTreeStr(t *testing.T) {
	if (nodefindergo.GetCleanTreeStr(rawTree) != cleanTree) {
		t.Error("GetCleanTreeStr (nodefindergo.go): Clean tree string failed")
	}
}

func Test_GetRightIndexOfName(t *testing.T) {
	expectedRightIndexOfName := 15

	result := nodefindergo.GetRightIndexOfName(cleanTree, nameA)

	if (result != expectedRightIndexOfName) {
		t.Error("GetRightIndexOfName (nodefindergo.go):\n" +
			fmt.Sprintf("Get right index of name failed (result: %d)", result))
	}
}

func Test_GetInsertionList(t *testing.T) {
	expectedInsertionList := []int{18, 19, 20, 27}

	result := nodefindergo.GetInsertionList(cleanTree, nameA)

	if (!reflect.DeepEqual(result, expectedInsertionList)) {
		t.Error("GetInsertionList (nodefindergo.go):\n" +
			fmt.Sprintf("Failed to get insertion list (result: %v)", result))
	}
}

func Test_GetIndexOfTMRCA(t *testing.T) {
	expectedIndexOfTMRCA := 19

	result := nodefindergo.GetIndexOfTMRCA(cleanTree, nameA, nameB)

	if (result != expectedIndexOfTMRCA) {
		t.Error("GetIndexOfTMRCA (nodefindergo.go):\n" +
			fmt.Sprintf("result: %v, expect: %v", result, expectedIndexOfTMRCA))
	}
}

func Test_SingleCalibration(t *testing.T) {
	expectedTreeWithCali := "((a,((b,c),(ddd,e))>0.1<0.2),(f,g));"

	result := nodefindergo.SingleCalibration(cleanTree, nameA, nameB, caliInfo)

	if (result != expectedTreeWithCali) {
		t.Error("SingleCalibration (nodefindergo.go):\n" +
			fmt.Sprintf("result: %v, expected: %v", result, expectedTreeWithCali))
	}
}

func Test_AddSingleBranchLabel(t *testing.T) {
	expectedTreeWithBranchLabel := "((a,((b,c),(ddd@0.3,e))),(f,g));"

	result := nodefindergo.AddSingleBranchLabel(cleanTree, nameA, branchLabel)

	if (!reflect.DeepEqual(result, expectedTreeWithBranchLabel)) {
		t.Error("AddSingleBranchLabel (nodefindergo.go):\n" +
			fmt.Sprintf("result: %v, expected: %v", result, expectedTreeWithBranchLabel))
	}
}

func Test_MultipleCalibration(t *testing.T) {
	expectedTreeStrWithMultiCalis := "((a, ((b, c)>0.03<0.05, (ddd, e))>0.1<0.2), (f#3, g));"
	calibrations := []nodefindergo.Calibration{
		{
			ID: 0,
			CaliType: nodefindergo.CALI_OR_CLADE_LABEL_TYPE,
			NameA: "b",
			NameB: "ddd",
			CaliInfo: ">0.1<0.2",
			Description: "First calibration"},
		{
			ID: 1,
			CaliType: nodefindergo.CALI_OR_CLADE_LABEL_TYPE,
			NameA: "b",
			NameB: "c",
			CaliInfo: ">0.03<0.05",
			Description: "Second calibration"},
		{
			ID: 2,
			CaliType: nodefindergo.BRANCH_LABEL_TYPE,
			NameA: "f",
			NameB: "",
			CaliInfo: "#3",
			Description: "First branchLabel"},
	}

	result := nodefindergo.MultipleCalibration(cleanTree, calibrations)

	if !(reflect.DeepEqual(result, expectedTreeStrWithMultiCalis)) {
		t.Error("MultipleCalibration (nodefindergo.go):\n" +
			fmt.Sprintf("result: %v, expected: %v", result, expectedTreeStrWithMultiCalis))
	}
}

func Test_ParseConfig(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Error("Get working dir failed")
	}

	caliFilePath := path.Join(workingDir, "calibration.txt")

	expectedCalibrations := []nodefindergo.Calibration{
		{
			ID: 0,
			CaliType: nodefindergo.CALI_OR_CLADE_LABEL_TYPE,
			NameA: "b",
			NameB: "ddd",
			CaliInfo: ">0.1<0.2",
			Description: "Normal calibration or clade label."},
		{
			ID: 1,
			CaliType: nodefindergo.CALI_OR_CLADE_LABEL_TYPE,
			NameA: "b",
			NameB: "c",
			CaliInfo: "$5",
			Description: "Normal calibration or clade label."},
		{
			ID: 2,
			CaliType: nodefindergo.BRANCH_LABEL_TYPE,
			NameA: "f",
			NameB: "",
			CaliInfo: "#3",
			Description: "Branch label description"},
	}

	result := nodefindergo.ParseConfig(caliFilePath)
	if (!reflect.DeepEqual(result, expectedCalibrations)) {
		t.Error("ParseConfig (nodefinder.go):\n" +
			fmt.Sprintf("result: %v,\nexpected: %v", result, expectedCalibrations))
	}
}
