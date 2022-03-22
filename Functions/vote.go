package Functions

import "strings"

func VerificaVoto(texto string) bool {
	var retorno bool = false
	var voto [1]string = [1]string{
		"voto n"}
	for i := 0; i < len(voto); i++ {
		if strings.Contains(texto, voto[i]) {
			retorno = true
		}
	}
	return retorno
}

func AnalisaVoto(ementaJulgado string, char int) string {
	texto := SelecionaUltimosChar(ementaJulgado, char)
	var tipo string
	if VerificaReexame(texto, char-24) {
		tipo = "Reexame Necessário"
	} else if VerificaDiligencia(texto, char-24) {
		tipo = "Convertido em Diligência"
	} else if VerificaPrejudicado(texto, char-24) {
		tipo = "Prejudicado"
	} else if VerificaParcial(texto, char-24) {
		tipo = "Parcial Procedente"
	} else if VerificaImprocedente(texto, char-24) {
		tipo = "Improcedente"
	} else if VerificaProcedente(texto, char-24) {
		tipo = "Procedente"
	} else {
		tipo = "Não Mapeado"
	}
	return tipo
}
