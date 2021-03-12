package utils

import (
	"testing"
)

type TestDataItem struct {
	input    string
	output   string
	hasError bool
}

func TestFilenameTrimmedSuffix(t *testing.T) {
	dataItems := []TestDataItem{
		{"testdata/file1.csv", "testdata/file1", false},
		{"testdata/file2.csv", "testdata/file2", false},
		{"testdata/file1", "testdata/file1", false},
	}

	for _, item := range dataItems {
		result := FilenameTrimmedSuffix(item.input)

		if result != item.output {
			t.Errorf("FilenameTrimmedSuffix() with args %v FAILED, expected %v but got value '%v'", item.input, item.output, result)
		} else {
			t.Logf("FilenameTrimmedSuffix() with args %v PASSED, expected %v and got value '%v'", item.input, item.output, result)
		}
	}
}

func TestCompareStringArrays(t *testing.T) {
	data := []string{"BRC", "SCL", "ORL", "CDG"}
	outputCorrect := []string{"ORL", "CDG", "BRC", "SCL"}
	outputWrong1 := []string{"GRU", "CDG", "BRC", "SCL"}
	outputWrong2 := []string{}

	if !CompareStringArrays(outputWrong1, data) {
		t.Logf("CompareStringArrays() with different arrays PASSED, expected %v and got %v", data, outputWrong1)
	} else {
		t.Errorf("CompareStringArrays() with different arrays FAILED, expected %v but got %v", data, outputWrong1)
	}

	if !CompareStringArrays(outputWrong2, data) {
		t.Logf("CompareStringArrays() with empty string array PASSED, expected %v and got %v", data, outputWrong2)
	} else {
		t.Errorf("CompareStringArrays() with empty string array FAILED, expected %v but got %v", data, outputWrong2)
	}

	if !CompareStringArrays(data, outputWrong2) {
		t.Logf("CompareStringArrays() with empty string array PASSED, expected %v and got %v", outputWrong2, data)
	} else {
		t.Errorf("CompareStringArrays() with empty string array FAILED, expected %v but got %v", outputWrong2, data)
	}

	if CompareStringArrays(outputCorrect, data) {
		t.Logf("CompareStringArrays() PASSED, expected %v and got %v", data, outputCorrect)
	} else {
		t.Errorf("CompareStringArrays() should have passed but FAILED, expected %v but got %v", data, outputCorrect)
	}
}

func TestTrimAndUpper(t *testing.T) {
	dataItems := []TestDataItem{
		{"BRC ", "BRC", true},
		{" ORL", "ORL", true},
		{" cdg ", "CDG", true},
	}

	for _, item := range dataItems {
		result := TrimAndUpper(item.input)

		if result != item.output {
			t.Errorf("TestTrimAndUpper() with args %v FAILED, expected '%v' but got value '%v'", item.input, item.output, result)
		} else {
			t.Logf("TestTrimAndUpper() with args %v PASSED, expected '%v' and got value '%v'", item.input, item.output, result)
		}
	}
}
