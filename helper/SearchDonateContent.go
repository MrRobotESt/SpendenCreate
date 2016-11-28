package helper

import (
	st "../structs"
	"github.com/Sirupsen/logrus"
)

// Search the spender in the DonateData and get his data.

func SearchDonateContent(name string, donateData []st.DonateData) []st.DonateData {

	var spender []st.DonateData

	logrus.Infoln("CheckDonateData")
	for x := range donateData {

		if name == donateData[x].Namen {

			spender = append(spender, st.DonateData{Namen: donateData[x].Namen, Betrag: donateData[x].Betrag, Buchungsdatum: donateData[x].Buchungsdatum})

		}
	}

	return spender
}
