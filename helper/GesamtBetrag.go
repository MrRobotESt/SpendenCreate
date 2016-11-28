package helper

import (
	st "../structs"
	"strconv"
	"strings"
)

// Create the summ of donations a person have made. (in Float!)

func GesamtBetrag(donatePersonData []st.DonateData) string {

	summe := 0.0
	for x := range donatePersonData {
		s, _ := strconv.ParseFloat(donatePersonData[x].Betrag, 2)
		summe += s
	}

	summeStringPoint := strconv.FormatFloat(summe, 'f', -1, 32)
	summeStringKomma := strings.Replace(summeStringPoint, ".", ",", 1)
	return summeStringKomma
}
