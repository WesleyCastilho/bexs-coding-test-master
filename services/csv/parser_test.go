package csvparser

import (
	"desafio-banco-bexs/services/utils"
	"os"
	"testing"
)

type TestDataItem struct {
	inputPath  string
	outputData [][]string
}

func TestReadCsv(t *testing.T) {
	pathFile1 := "testdata/file1.csv"
	outputFile1 := [][]string{
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
	}
	pathFile2 := "testdata/file2.csv"
	outputFile2 := [][]string{
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
	}

	dataItems := []TestDataItem{
		{pathFile1, outputFile1},
		{pathFile2, outputFile2},
	}

	for _, item := range dataItems {
		err, result := Read(item.inputPath)

		if err != nil {
			t.Errorf("csvparser.Read() do caminho %v FALHOU, esperava %v mas  obtivemos um erro '%v'", item.inputPath, item.outputData, err)
		} else if len(result) != 2 {
			t.Errorf("csvparser.Read() do caminho %v FALHOU, esperava %v mas obtivemos o valor '%v'", item.inputPath, item.outputData, result)
		} else {
			t.Logf("csvparser.Read() do caminho %v PASSOU, esperava %v e  obtivemos '%v'", item.inputPath, item.outputData, result)
		}
	}
}

func BeforeTestWriteCsv(filePaths []string) []string {
	result := make([]string, len(filePaths))

	for index, item := range filePaths {
		newPathName := utils.FilenameTrimmedSuffix(item) + "_testwrite.csv"
		err := utils.Copy(item, newPathName)
		if err != nil {
			result[index] = item
		} else {
			result[index] = newPathName
		}
	}

	return result
}

func AfterTestWriteCsv(filePaths []string) {
	for _, item := range filePaths {
		os.Remove(item)
	}
}
func TestWriteCsv(t *testing.T) {
	inputPathFile1 := "testdata/file1.csv"
	outputFile1 := [][]string{
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
	}
	inputPathFile2 := "testdata/file2.csv"
	outputFile2 := [][]string{
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
	}

	outputPaths := BeforeTestWriteCsv([]string{inputPathFile1, inputPathFile2})

	dataReadItems := []TestDataItem{
		{inputPathFile1, outputFile1},
		{inputPathFile2, outputFile2},
	}
	dataWriteItems := []TestDataItem{
		{outputPaths[0], outputFile1},
		{outputPaths[1], outputFile2},
	}

	for index, item := range dataReadItems {
		errRead, readResult := Read(item.inputPath)

		if errRead != nil {
			t.Errorf("csvparser.Write() do caminho %v FALHOU, Read() não teve sucesso, como esperava %v mas  obtivemos um erro '%v'", item.inputPath, item.outputData, errRead)
		} else {
			errWrite, writeResult := Write(dataWriteItems[index].inputPath, readResult)
			if errWrite != nil {
				t.Errorf("csvparser.Write() do caminho %v FALHOU, esperava %v mas obteve erro '%v'", dataWriteItems[index], item.outputData, errWrite)
			} else if len(writeResult) != 4 {
				t.Errorf("csvparser.Write() do caminho %v FALHOU, esperava %v mas obtivemos o valor '%v'", dataWriteItems[index], item.outputData, writeResult)
			} else {
				t.Logf("csvparser.Write() do caminho %v PASSOU, esperava %v e  obtivemos '%v'", dataWriteItems[index], item.outputData, writeResult)
			}
		}
	}

	AfterTestWriteCsv(outputPaths)
}
func TestCreateWrite(t *testing.T) {
	inputPathFile1 := "testdata/createwrite_test1.csv"
	outputFile1 := [][]string{
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
		{"GRU", "BRC", "20"},
		{"BRC", "SCL", "5"},
	}
	inputPathFile2 := "testdata/createwrite_test1.csv"
	outputFile2 := [][]string{
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
		{"GRU", "CDG", "75"},
		{"SCL", "ORL", "20"},
	}

	dataWriteItems := []TestDataItem{
		{inputPathFile1, outputFile1},
		{inputPathFile2, outputFile2},
	}

	for _, item := range dataWriteItems {
		err, result := CreateWrite(item.inputPath, item.outputData)

		if err != nil {
			t.Errorf("csvparser.CreateWrite() do caminho %v FALHOU, como esperava %v mas obtivemos um erro '%v'", item.inputPath, item.outputData, err)
		} else if len(result) != 4 {
			t.Errorf("csvparser.CreateWrite() do caminho %v FALHOU, esperava %v mas obtivemos o valor '%v'", item.inputPath, item.outputData, result)
		} else {
			t.Logf("csvparser.CreateWrite() do caminho %v PASSOU, esperava %v e obtivemos '%v'", item.inputPath, item.outputData, result)
		}
	}

	AfterTestWriteCsv([]string{inputPathFile1, inputPathFile2})
}

func TestEmptyDataCreateWrite(t *testing.T) {
	inputPathFile1 := "testdata/createwrite_test1.csv"
	inputPathFile2 := "testdata/createwrite_test1.csv"

	dataWriteItems := []TestDataItem{
		{inputPathFile1, nil},
		{inputPathFile2, nil},
	}

	for _, item := range dataWriteItems {
		err, result := CreateWrite(item.inputPath, item.outputData)

		if err != nil && err.Error() == "Não há dados para serem gravados." && len(result) == 0 {
			t.Logf("csvparser.CreateWrite() do caminho %v PASSOU, esperava um erro 'Não há dados para serem gravados.' e  obtivemos um erro '%v'", item.inputPath, err.Error())
		} else {
			t.Errorf("csvparser.CreateWrite() do caminho %v FALHOU, esperava um erro 'Não há dados para serem gravados.' mas não obtive erro", item.inputPath)
		}
	}

	AfterTestWriteCsv([]string{inputPathFile1, inputPathFile2})
}

func TestEmptyDataWrite(t *testing.T) {
	inputPathFile1 := "testdata/file1.csv"
	inputPathFile2 := "testdata/file2.csv"

	outputPaths := BeforeTestWriteCsv([]string{inputPathFile1, inputPathFile2})
	dataWriteItems := []TestDataItem{
		{outputPaths[0], nil},
		{outputPaths[1], nil},
	}

	for _, item := range dataWriteItems {
		errWrite, writeResult := Write(item.inputPath, nil)

		if errWrite == nil && errWrite.Error() == "Sem dados para serem gravados." && len(writeResult) == 0 {
			t.Logf("csvparser.Write() do caminho %v PASSOU, esperava um erro 'Sem dados para serem gravados.' e  obtivemos um erro '%v'", item.inputPath, errWrite.Error())
		} else {
			t.Errorf("csvparser.Write() do caminho %v FALHOU, esperava um erro 'Sem dados para serem gravados.' mas não obtive erro", item.inputPath)
		}
	}

	AfterTestWriteCsv(outputPaths)
}

func TestEmptyPathCreateWrite(t *testing.T) {
	dataWriteItems := []TestDataItem{
		{"", [][]string{}},
		{"", [][]string{}},
	}

	for _, item := range dataWriteItems {
		err, result := CreateWrite(item.inputPath, item.outputData)

		if err != nil && err.Error() == "o Path está vazio, por favor tente novamente" && len(result) == 0 {
			t.Logf("csvparser.CreateWrite() do caminho %v PASSOU, esperava um erro 'o Path está vazio, por favor tente novamente' e  obtivemos um erro '%v'", item.inputPath, err.Error())
		} else {
			t.Errorf("csvparser.CreateWrite() do caminho %v FALHOU, esperava um erro 'o Path está vazio, por favor tente novamente' mas não obtive erro", item.inputPath)
		}
	}
}

func TestEmptyPathWrite(t *testing.T) {
	dataWriteItems := []TestDataItem{
		{"", [][]string{}},
		{"", [][]string{}},
	}

	for _, item := range dataWriteItems {
		errWrite, writeResult := Write(item.inputPath, item.outputData)

		if errWrite != nil && errWrite.Error() == "o Path está vazio, por favor tente novamente" && len(writeResult) == 0 {
			t.Logf("csvparser.Write() do caminho %v PASSOU, esperava um erro 'o Path está vazio, por favor tente novamente' e  obtivemos um erro '%v'", item.inputPath, errWrite.Error())
		} else {
			t.Errorf("csvparser.Write() do caminho %v FALHOU, esperava um erro 'o Path está vazio, por favor tente novamente' mas não obtive erro", item.inputPath)
		}
	}
}
