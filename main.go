package main

import (
	"fmt"
	"github.com/Darklabel91/Decision_Classifier/CSV"
	"github.com/Darklabel91/Decision_Classifier/Classifier"
)

func main() {

	//Example with a single summary decision
	summary := " CONDOMÍNIO. COBRANÇA. COMPROMISSO DE COMPRAE VENDA. ILEGITIMIDADE PASSIVA DE PARTE. DÍVIDA\"PROPTER REM\". PREVALÊNCIA DO INTERESSE DACOLETIVIDADE CONDOMINIAL. PROPOSITURA DADEMANDA EM FACE DE QUEM FOR MAIS CONVENIENTE,OU SEJA, CONTRA QUEM PODERÁ CUMPRIR MAISPRONTAMENTE A OBRIGAÇÃO. Como o condomínio elegeu otitular do domínio (nome constante do álbum imobiliário), comoo responsável pelo débito referente ao imóvel em questão,apoiando-se nos termos do que dispõe a lei, e amparado ainda,em v. precedentes jurisprudenciais, nesse sentido, a sua opçãodeve ser respeitada, pois o interesse prevalente é o dacoletividade de receber os recursos para o pagamento dedespesas indispensáveis e inadiáveis, ressalvado a quem cumprira obrigação exercer o direito regressivo contra quem entendaresponsável. RECURSO PROVIDO. "
	id := "0"
	court := "nowhere"
	inferedDecision := Classifier.Decision_Classifier(summary, id, court)

	fmt.Println("Summary:", inferedDecision.Summary)
	fmt.Println("Text:", inferedDecision.Text)
	fmt.Println("Class:", inferedDecision.Class)
	fmt.Println("Id:", inferedDecision.Identifier)
	fmt.Println("Court", inferedDecision.Court)

	//Example with CSV reading [CSV must be composed by: {id, identifier, decision, court}]
	CSV.Loop(CSV.ReadCsvFile("CSV/Database/testData.csv"))
}
