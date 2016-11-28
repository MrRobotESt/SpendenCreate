package helper

import (
	st "../structs"
)

//Get the TimePeriod

func Period(dep []st.DonateData) string {

	period := dep[0].Buchungsdatum + " - " + dep[len(dep)-1].Buchungsdatum

	return period
}
