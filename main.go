package main

import (
	"fmt"
	"github.com/Darklabel91/Decision_Classifier/CSV"
	"github.com/Darklabel91/Decision_Classifier/Classifier"
	"github.com/Darklabel91/Decision_Classifier/Struct"
	"time"
)

var result []Struct.Infered_decision

func main() {

	//Example with a single summary decision
	summary := " CONDOMÍNIO. COBRANÇA. COMPROMISSO DE COMPRAE VENDA. ILEGITIMIDADE PASSIVA DE PARTE. DÍVIDA\"PROPTER REM\". PREVALÊNCIA DO INTERESSE DACOLETIVIDADE CONDOMINIAL. PROPOSITURA DADEMANDA EM FACE DE QUEM FOR MAIS CONVENIENTE,OU SEJA, CONTRA QUEM PODERÁ CUMPRIR MAISPRONTAMENTE A OBRIGAÇÃO. Como o condomínio elegeu otitular do domínio (nome constante do álbum imobiliário), comoo responsável pelo débito referente ao imóvel em questão,apoiando-se nos termos do que dispõe a lei, e amparado ainda,em v. precedentes jurisprudenciais, nesse sentido, a sua opçãodeve ser respeitada, pois o interesse prevalente é o dacoletividade de receber os recursos para o pagamento dedespesas indispensáveis e inadiáveis, ressalvado a quem cumprira obrigação exercer o direito regressivo contra quem entendaresponsável. RECURSO PROVIDO. "
	id := "0"
	court := "nowhere"
	fmt.Println(Classifier.Decision_Classifier(summary, id, court))

	//Example with CSV reading
	loop(CSV.ReadCsvFile("CSV/Database/decision_Short.csv"))
}

func loop(decisions []Struct.Raw_decision) {
	start := time.Now()

	for i := 0; i < len(decisions); i++ {
		retorno := Classifier.Decision_Classifier(decisions[i].Summary, decisions[i].Identifier, decisions[i].Court)
		result = append(result, retorno)
		CSV.DataComp(retorno)
	}

	finish := time.Since(start).Minutes()
	fmt.Println(len(decisions), " - Classified in ", finish, "minutes")

	CSV.ExportData(result)

}
