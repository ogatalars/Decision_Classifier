package Functions

import "strings"

func VerificaParcial(texto string, char int) bool {
	var retorno bool = false
	var parcial []string
	if char < 17 {
		parcial = append(parcial,
			"adequação parcial",
			"almente mantid",
			"arte conhecid",
			"cial provimento",
			"dência parcial",
			"em parte",
			"ialmente",
			"lmente",
			"mente",
			"nessa parte",
			"nesta parte",
			"neste tópico",
			"parcial",
			"parcial provimento",
			"parcialmente modificada",
			"parcialmente modificado",
			"parcialmente provid",
			"parcialmente",
			"parte acolhid",
			"parte apreciad",
			"parte chonhecid",
			"parte e provid",
			"parte provid",
			"parte, portanto",
			"provido, em parte",
			"provido em parte",
			"rcial provimento",
			"reformada parcialmente",
			"te provid")
	} else {
		parcial = append(parcial,
			"adequação parcial",
			"em parte",
			"nessa parte",
			"nesta parte",
			"parcialmente",
			"conhecido neste tópico",
			"em parte acolhid",
			"em parte apreciad",
			"em parte conhecid",
			"em parte e provid",
			"em parte provid",
			"em parte, portanto",
			"parcial provimento",
			"parcial provimento",
			"parcialmente mantid",
			"parcialmente modificada",
			"parcialmente modificado",
			"paricalmente provid",
			"parcialmente provid",
			"parte conhecid",
			"procedência parcial",
			"provido, em parte",
			"provido em parte",
			"reformada parcialmente")
	}
	for i := 0; i < len(parcial); i++ {
		if strings.Contains(texto, parcial[i]) {
			retorno = true
		}
	}
	return retorno
}
