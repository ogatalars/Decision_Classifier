package CSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Decision_Classifier/Error"
	"github.com/Darklabel91/Decision_Classifier/Struct"
	"os"
	"path/filepath"
)

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ExportCSV(fileName string, result []Struct.Infered_decision) {
	empData := [][]string{}

	empData = append(empData, []string{"Summary", "Text", "Class", "Identifier", "Court"})

	for i := 0; i < len(result); i++ {
		final := []string{result[i].Summary, result[i].Text, result[i].Class, result[i].Identifier, result[i].Court}
		empData = append(empData, final)
	}

	csvFile, _ := create("CSV/Files CSV/" + fileName + ".csv")
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range empData {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	err := csvFile.Close()
	Error.CheckError(err)
}
