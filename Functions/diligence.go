package Functions

import "strings"

func VerificaDiligencia(texto string, char int) bool {
	var retorno bool = false
	var diligencia []string
	if char < 17 {
		diligencia = append(diligencia,
			"conversão do julgamento em diligência",
			"convertido em diligência",
			"diligência")
	} else {
		diligencia = append(diligencia,
			"acórdão em diligência",
			"conversão do julgamento em diligência",
			"conversão do feito em diligência",
			"convertido em diligência",
			"decisão em diligência",
			"julgalmento em diligência",
			"sentença em diligência")
	}
	for i := 0; i < len(diligencia); i++ {
		if strings.Contains(texto, diligencia[i]) {
			retorno = true
		}

	}
	return retorno
}
