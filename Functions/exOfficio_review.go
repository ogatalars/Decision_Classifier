package Functions

import "strings"

func ExOfficioReview(text string, char int) bool {
	var ret bool = false
	var exOfficioReview []string
	if char < 17 {
		exOfficioReview = append(exOfficioReview,
			"reexame necessário", "xame necessário")
	} else {
		exOfficioReview = append(exOfficioReview,
			"reexame necessário")
	}
	for i := 0; i < len(exOfficioReview); i++ {
		if strings.Contains(text, exOfficioReview[i]) {
			ret = true
		}
	}
	return ret
}
