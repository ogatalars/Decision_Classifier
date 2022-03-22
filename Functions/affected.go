package Functions

import "strings"

func Affected(text string, char int) bool {
	var ret bool = false
	var affected []string
	if char < 17 {
		affected = append(affected,
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
		affected = append(affected,
			"extinção da punibilidade",
			"perda do objeto",
			"prejudicad",
			"prejudicado",
			"prejudicada",
			"perda superveniente do objeto",
			"prejudicando o exame")
	}
	for i := 0; i < len(affected); i++ {
		if strings.Contains(text, affected[i]) {
			ret = true
		}
	}
	return ret
}
