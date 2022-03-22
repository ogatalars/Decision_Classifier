package Functions

import "strings"

func VerificaNumero(texto string) bool {
	var retorno bool = false
	var letras [28]string = [28]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "x", "w", "y", "z", "/", ")"}
	var numero [10]string = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for i := 0; i < len(numero); i++ {
		if strings.Contains(texto, numero[i]) {
			retorno = true
		}
	}
	if retorno {
		for a := 0; a < len(letras); a++ {
			if strings.Contains(texto, "eros") {
				retorno = true
			} else {
				if strings.Contains(texto, letras[a]) {
					retorno = false
				}
			}
		}
	}
	return retorno
}

func AnalisaNumero(ementa string, char int) string {
	var retorno string
	tamanho := len(ementa)
	texto := SelecionaUltimosChar(ementa, tamanho*2)
	if tamanho < char {
		retorno = "Sem Informação"
	} else {
		if VerificaReexame(texto, char-26) {
			retorno = "Reexame Necessário"
		} else if VerificaDiligencia(texto, char-26) {
			retorno = "Convertido em Diligência"
		} else if VerificaPrejudicado(texto, char-26) {
			retorno = "Prejudicado"
		} else if VerificaParcial(texto, char-26) {
			retorno = "Parcial Procedente"
		} else if VerificaImprocedente(texto, char-26) {
			retorno = "Improcedente"
		} else if VerificaProcedente(texto, char-26) {
			retorno = "Procedente"
		} else {
			retorno = "Não Mapeado"
		}
	}
	return retorno
}
