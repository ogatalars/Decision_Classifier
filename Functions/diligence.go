package Functions

import "strings"

func Diligence(text string, char int) bool {
	var ret bool = false
	var diligence []string
	if char < 17 {
		diligence = append(diligence,
			"conversão do julgamento em diligência",
			"convertido em diligência",
			"diligência")
	} else {
		diligence = append(diligence,
			"acórdão em diligência",
			"conversão do julgamento em diligência",
			"conversão do feito em diligência",
			"convertido em diligência",
			"decisão em diligência",
			"julgalmento em diligência",
			"sentença em diligência")
	}
	for i := 0; i < len(diligence); i++ {
		if strings.Contains(text, diligence[i]) {
			ret = true
		}

	}
	return ret
}
