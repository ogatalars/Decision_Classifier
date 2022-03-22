package Functions

import "strings"

func VerificaPrejudicado(texto string, char int) bool {
	var retorno bool = false
	var prejudicado []string
	if char < 17 {
		prejudicado = append(prejudicado,
			"dicado",
			"extinção da punibilidade",
			"icado",
			"perda do objeto",
			"prejudicad",
			"prejudicado",
			"prejudicada",
			"perda superveniente do objeto",
			"prejudicando o exame")
	} else {
		prejudicado = append(prejudicado,
			"extinção da punibilidade",
			"perda do objeto",
			"prejudicad",
			"prejudicado",
			"prejudicada",
			"perda superveniente do objeto",
			"prejudicando o exame")
	}
	for i := 0; i < len(prejudicado); i++ {
		if strings.Contains(texto, prejudicado[i]) {
			retorno = true
		}
	}
	return retorno
}
