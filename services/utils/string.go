package utils

import (
	"path/filepath"
	"sort"
	"strings"
)

func FilenameTrimmedSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func CompareStringArrays(arr1 []string, arr2 []string) bool {
	if len(arr2) == 0 || len(arr1) == 0 {
		return false
	}

	sort.Strings(arr1)
	sort.Strings(arr2)

	for _, item1 := range arr1 {
		i := sort.SearchStrings(arr2, item1)
		if i >= len(arr2) || arr2[i] != item1 {
			return false
		}
	}

	return true
}

func TrimAndUpper(str string) string {
	return strings.Trim(strings.ToUpper(str), "\t \n")
}
