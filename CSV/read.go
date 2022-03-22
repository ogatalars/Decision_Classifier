package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Decision_Classifier/Error"
	"github.com/Darklabel91/Decision_Classifier/Struct"
	"os"
)

func ReadCsvFile(filePath string) []Struct.Raw_decision {
	var data2 []Struct.Raw_decision

	csvFile, err := os.Open(filePath)
	Error.CheckError(err)

	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {

		}
	}(csvFile)

	csvLines, err0 := csv.NewReader(csvFile).ReadAll()
	Error.CheckError(err0)

	for _, line := range csvLines {

		emp := Struct.Raw_decision{
			Summary:    line[2],
			Identifier: line[1],
			Court:      line[3],
		}

		data2 = append(data2, emp)
	}
	return data2
}
