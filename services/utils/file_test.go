package utils

import (
	"testing"
)

type TestCopyDataItem struct {
	src      string
	dst      string
	hasError bool
}

type TestFileExistsDataItem struct {
	input    string
	output   bool
	hasError bool
}

func TestCopy(t *testing.T) {
	dataFailItems := []TestCopyDataItem{
		{"_testread.csv", "_testwrite.csv", false},
		{"../csv/testdata/file1.csv", "/etc/testtt", false},
	}

	dataSuccessItems := []TestCopyDataItem{
		{"../csv/testdata/file1.csv", "testdata/_testwrite.csv", true},
	}

	for _, item := range dataFailItems {
		err := Copy(item.src, item.dst)

		if err != nil {
			t.Logf("Copy() with args %v->%v PASSED, expected an error and got error '%v'", item.src, item.dst, err)
		} else {
			t.Errorf("Copy() with args %v->%v FAILED, expected an error but got no error", item.src, item.dst)
		}
	}

	for _, item := range dataSuccessItems {
		err := Copy(item.src, item.dst)

		if err == nil {
			t.Logf("Copy() with args %v->%v PASSED, expected no error and got no error", item.src, item.dst)
		} else {
			t.Errorf("Copy() with args %v->%v FAILED, expected no error but got an error '%v'", item.src, item.dst, err)
		}
	}

}

func TestFileExists(t *testing.T) {
	dataItems := []TestFileExistsDataItem{
		{"file.go", true, true},
		{"file123.go", false, false},
		{"../utils", false, false},
	}

	for _, item := range dataItems {
		result := FileExists(item.input)

		if result != item.output {
			t.Errorf("FileExists() with args %v FAILED, expected '%v' but got value '%v'", item.input, item.output, result)
		} else {
			t.Logf("FileExists() with args %v PASSED, expected '%v' and got value '%v'", item.input, item.output, result)
		}
	}
}
