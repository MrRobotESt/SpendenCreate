package helper


import (
	"fmt"
	st "../structs"
)


func PersonExistsDepositChecker(dep []st.DonateData) bool {
	fmt.Println(len(dep))
	if len(dep) > 0 {
		return true
	}

	return false
}

func NamensOnlyInDeposit(dep []st.DonateData, pers []st.AdressData) {

	var namens string

	for x := range dep {

		for y := range pers {

			if dep[x].Namen == pers[y].Namen  {
				continue
			} else {
				namens += dep[x].Namen + "\n"
			}

		}


	}

	if namens != "" {
		WriteToFile("OnlyInDeposit", namens)
	}



}