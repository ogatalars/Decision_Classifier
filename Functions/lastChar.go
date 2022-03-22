package Functions

import "strings"

func SelecionaUltimosChar(ementa string, char int) string {
	var textoFinal string
	if len(ementa) >= char {
		textoFinal = strings.ToLower(ementa[len(ementa)-char:])
	} else {
		textoFinal = "Ementa pequena"
	}
	return textoFinal
}
