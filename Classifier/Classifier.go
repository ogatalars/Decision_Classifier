package Classifier

import (
	"Decision_Classifier/Functions"
	"Decision_Classifier/Struct"
)

func Decision_Classifier(ementaJulgado string, identifier string, court string) Struct.Infered_decision {

	retorno := classificaJulgado(ementaJulgado, identifier, court, 16)

	if retorno.Classificcao == "Não Mapeado" {
		retorno = classificaJulgado(ementaJulgado, identifier, court, 70)
		if retorno.Classificcao == "Não Mapeado" || retorno.Classificcao == "Sem informaçao" {
			retorno = classificaJulgado(ementaJulgado, identifier, court, len(ementaJulgado)/2)
			if retorno.Classificcao == "Não Mapeado" {
				retorno = classificaJulgado(ementaJulgado, identifier, court, len(ementaJulgado))
			}
		}
	}

	return retorno

}

func classificaJulgado(ementaJulgado string, identifier string, court string, char int) Struct.Infered_decision {
	var tipo string
	var texto = Functions.SelecionaUltimosChar(ementaJulgado, char)

	if Functions.VerificaReexame(texto, char) {
		tipo = "Reexame Necessário"
	} else if Functions.VerificaDiligencia(texto, char) {
		tipo = "Convertido em Diligência"
	} else if Functions.VerificaPrejudicado(texto, char) {
		tipo = "Prejudicado"
	} else if Functions.VerificaParcial(texto, char) {
		tipo = "Parcial Procedente"
	} else if Functions.VerificaImprocedente(texto, char) {
		tipo = "Improcedente"
	} else if Functions.VerificaProcedente(texto, char) {
		tipo = "Procedente"
	} else if Functions.VerificaNumero(texto) {
		tipo = Functions.AnalisaNumero(ementaJulgado, char+26)
	} else if Functions.VerificaVoto(texto) {
		tipo = Functions.AnalisaVoto(ementaJulgado, char+24)
	} else if texto == "Ementa pequena" {
		tipo = "Sem Informação"
	} else {
		tipo = "Não Mapeado"
	}

	return Struct.Infered_decision{
		Identifier:   identifier,
		Classificcao: tipo,
		Court:        court,
	}
}
