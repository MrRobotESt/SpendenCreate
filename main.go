package main

import (
	hp "./helper"
	sp "./spendenCreator"
	st "./structs"
	"flag"
	"sort"
	"strings"
)

/* Control Sequenz */
var csvConverter = flag.Bool("csvconvert", false, "csv has to be convert in german iso8859-1 Chars")
var startCreator = flag.Bool("start", false, "start the pdf creator")
var checkexistence = flag.Bool("existence", false, "check the existence of Person in DepositData")

//Struct for sorting DonateData
type Deposit []st.DonateData

func main() {

	flag.Parse()
	used := func(a *flag.Flag) {

		switch {

		//Only convert the CSV Files into UTF8
		case a.Name == "csvconvert":
			hp.CSVtoUTF8("adressen")
			hp.CSVtoUTF8("buchung")

		//Start and create the PDFs
		case a.Name == "start":

			pC := sp.PageContentBuilder()
			adressSlice := hp.GetNameAdressDataFromCSV()
			donateData := hp.GetDonateData()

			for x := range adressSlice {

				donatePersonData := hp.SearchDonateContent(adressSlice[x].Namen, donateData)
				summeString := hp.GesamtBetrag(donatePersonData)
				sort.Sort(Deposit(donatePersonData))
				check := hp.PersonExistsDepositChecker(donatePersonData)
				if check == true {
					period := hp.Period(donatePersonData)
					sp.PageBuilder(pC, adressSlice[x], donatePersonData, summeString, period)
				} else {
					//Automatic Check of Person from Adress.csv exist in Buchungs.csv!
					hp.WriteToFile("NoDeposit", adressSlice[x].Namen)
				}

			}

		//Check the existence of a Person from BUCHUNG.csv in Adress.csv!
		case a.Name == "existence":

			adressSlice := hp.GetNameAdressDataFromCSV()
			donateData := hp.GetDonateData()
			hp.NamensOnlyInDeposit(donateData, adressSlice)

		}

	}
	flag.Visit(used)
}

/* Help Funcitons */

//Needed Sort Libraryfunctions for overriding!

func (a Deposit) Len() int      { return len(a) }
func (a Deposit) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Deposit) Less(i, j int) bool {
	datumSliceA := strings.Split(a[i].Buchungsdatum, ".")
	datumSliceB := strings.Split(a[j].Buchungsdatum, ".")

	if datumSliceA[2] < datumSliceB[2] {
		return true
	}

	return datumSliceA[0] <= datumSliceB[0] && datumSliceA[1] <= datumSliceB[1]
}
