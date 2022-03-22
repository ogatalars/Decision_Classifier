package Functions

import "strings"

func VerificaReexame(texto string, char int) bool {
	var retorno bool = false
	var reexame []string
	if char < 17 {
		reexame = append(reexame,
			"reexame necessário", "xame necessário")
	} else {
		reexame = append(reexame,
			"reexame necessário")
	}
	for i := 0; i < len(reexame); i++ {
		if strings.Contains(texto, reexame[i]) {
			retorno = true
		}
	}
	return retorno
}
