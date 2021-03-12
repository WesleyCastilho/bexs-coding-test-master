package csvparser

import (
	"desafio-banco-bexs/services/utils"
	"encoding/csv"
	"errors"
	"os"
)

func Read(path string) (error, [][]string) {
	csvfile, err := os.Open(path)
	if err != nil {
		return err, nil
	}

	r := csv.NewReader(csvfile)

	records, _ := r.ReadAll()
	return nil, records
}

func CreateWrite(path string, csvData [][]string) (error, [][]string) {
	if path == "" {
		return errors.New("o Path está vazio, por favor tente novamente"), nil
	}

	if csvData == nil {
		return errors.New("Não há dados para serem gravados."), nil
	}

	csvfile, err := os.Create(path)
	if err != nil {
		return errors.New("Não foi possível criar o arquivo para ser gravado!"), nil
	}

	writer := csv.NewWriter(csvfile)

	err = writer.WriteAll(csvData)
	if err != nil {
		return err, nil
	}

	return nil, csvData
}

func Write(path string, csvData [][]string) (error, [][]string) {
	if csvData == nil {
		return errors.New("Não há dado a ser gravado."), nil
	}

	if path == "" {
		return errors.New("o Path está vazio, por favor tente novamente"), nil
	}

	dataVals := [][]string{}
	var dataFile [][]string

	if !utils.FileExists(path) {
		_, err := os.Create(path)
		if err != nil {
			return errors.New("Poxa, não podemos gravar um arquivo inexistente!"), nil
		}
	} else {
		err, data := Read(path)
		if err != nil {
			return errors.New("Não conseguimos abrir o arquivo para gravação!"), nil
		}
		dataFile = data
	}

	if dataFile == nil {
		dataFile = dataVals
	} else {
		for _, item := range csvData {
			dataFile = append(dataFile, item)
		}
	}

	csvfile, err := os.Create(path)
	if err != nil {
		return err, nil
	}

	w := csv.NewWriter(csvfile)

	err = w.WriteAll(dataFile)
	if err != nil {
		return err, nil
	}

	return nil, dataFile
}
