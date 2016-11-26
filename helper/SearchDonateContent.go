package helper

import (
	"fmt"
st "../structs"


)

// Search the spender in the DonateData and get his data.

func SearchDonateContent(name string, donateData []st.DonateData) []st.DonateData{

	var spender []st.DonateData


	for x := range donateData {
		fmt.Println("Holla")
		if name == donateData[x].Namen {
			spender = append(spender, st.DonateData{Namen: donateData[x].Namen, Betrag: donateData[x].Betrag, Buchungsdatum: donateData[x].Buchungsdatum})
			fmt.Println("HLLO")
		}
	}

	return spender
}