package main

import (
	sp "./spendenCreator"
	"flag"
	"fmt"
	"sort"
	"strings"
	hp "./helper"
	st "./structs"
)



/* Control Sequenz */
var csvConverter = flag.Bool("csvconvert",false,"csv has to be convert in german iso8859-1 Chars" )
var startCreator = flag.Bool("start",false,"start the pdf creator" )
var checkexistence = flag.Bool("existence", false, "check the existence of Person in DepositData")


//Struct for sorting DonateData
type Deposit []st.DonateData

func main() {

	flag.Parse()


fmt.Println("OK")

	switch{

	case *csvConverter == true:
		fmt.Println("TEST")




	case *startCreator == true:


		hp.CSVtoUTF8("adressen")
		hp.CSVtoUTF8("buchung")
		pC := sp.PageContentBuilder()

		adressSlice := hp.GetNameAdressDataFromCSV()
		donateData := hp.GetDonateData()
		fmt.Println(donateData)
		for x := range adressSlice {
			fmt.Println(adressSlice[x].Namen)
			donatePersonData := hp.SearchDonateContent(adressSlice[x].Namen, donateData)
			summeString := hp.GesamtBetrag(donatePersonData)
			fmt.Println(donatePersonData)
			sort.Sort(Deposit(donatePersonData))

			fmt.Println(len(donatePersonData))
			check := hp.PersonExistsDepositChecker(donatePersonData)
			if check == true {
				period := hp.Period(donatePersonData)
				sp.PageBuilder(pC, adressSlice[x], donatePersonData, summeString, period)
			} else {

				hp.WriteToFile("NoDeposit", adressSlice[x].Namen)
			}

		}

		hp.NamensOnlyInDeposit(donateData, adressSlice)



	}

}

/* Help Funcitons */


//Needed Sort Libraryfunctions for overriding!

func (a Deposit) Len() int           { return len(a) }
func (a  Deposit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a  Deposit) Less(i, j int) bool {
	datumSliceA := strings.Split(a[i].Buchungsdatum,".")
	datumSliceB := strings.Split(a[j].Buchungsdatum,".")
	fmt.Println("*****************")
	fmt.Println(a[i].Buchungsdatum)
	fmt.Println(a[j].Buchungsdatum)

	if datumSliceA[2] < datumSliceB[2]{
		return true
	}

	return datumSliceA[0] <= datumSliceB[0] &&  datumSliceA[1] <= datumSliceB[1]
}