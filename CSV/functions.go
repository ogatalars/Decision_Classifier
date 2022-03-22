package CSV

import (
	"fmt"
	"github.com/Darklabel91/Decision_Classifier/Classifier"
	"github.com/Darklabel91/Decision_Classifier/Struct"
	"time"
)

var exofficioReview []Struct.Infered_decision
var diligence []Struct.Infered_decision
var affected []Struct.Infered_decision
var partial []Struct.Infered_decision
var groundless []Struct.Infered_decision
var hasGrounds []Struct.Infered_decision
var noSummary []Struct.Infered_decision
var notMapped []Struct.Infered_decision

func Loop(decisions []Struct.Raw_decision) {
	var result []Struct.Infered_decision
	start := time.Now()

	for i := 0; i < len(decisions); i++ {
		retorno := Classifier.Decision_Classifier(decisions[i].Summary, decisions[i].Identifier, decisions[i].Court)
		result = append(result, retorno)
		dataComp(retorno)
	}

	finish := time.Since(start).Minutes()
	fmt.Println(len(decisions), " - Classified in ", finish, "minutes")

	exportData(result)

}

func dataComp(decision Struct.Infered_decision) {
	if decision.Class == "Reexame Necessário" {
		exofficioReview = append(exofficioReview, decision)
	} else if decision.Class == "Convertido em Diligência" {
		diligence = append(diligence, decision)
	} else if decision.Class == "Prejudicado" {
		affected = append(affected, decision)
	} else if decision.Class == "Parcial Procedente" {
		partial = append(partial, decision)
	} else if decision.Class == "Improcedente" {
		groundless = append(groundless, decision)
	} else if decision.Class == "Procedente" {
		hasGrounds = append(hasGrounds, decision)
	} else if decision.Class == "Sem Informação" {
		noSummary = append(noSummary, decision)
	} else if decision.Class == "Não Mapeado" {
		notMapped = append(notMapped, decision)
	}

}
func exportData(result []Struct.Infered_decision) {
	start := time.Now()
	ExportCSV("exofficioReview", exofficioReview)
	ExportCSV("diligence", diligence)
	ExportCSV("affected", affected)
	ExportCSV("partial", partial)
	ExportCSV("groundless", groundless)
	ExportCSV("hasGrounds", hasGrounds)
	ExportCSV("noSummary", noSummary)
	ExportCSV("notMapped", notMapped)
	ExportCSV("Total_Classification", result)

	finish := time.Since(start).Minutes()
	fmt.Println("9 csv created in ", finish, "minutes")

}
