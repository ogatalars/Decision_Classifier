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

	defer csvFile.Close()

	csvLines := csv.NewReader(csvFile)
	csvLines.Comma = ';'
	csvData, err1 := csvLines.ReadAll()
	Error.CheckError(err1)

	for _, line := range csvData {

		emp := Struct.Raw_decision{
			Summary:    line[2],
			Identifier: line[1],
			Court:      line[3],
		}

		data2 = append(data2, emp)
	}
	return data2
}
