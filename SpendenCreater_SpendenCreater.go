package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"encoding/csv"
	"log"
	"time"
	"github.com/Sirupsen/logrus"

	"strconv"
	"sort"
)

type AdressData struct {
	Namen string
	Firma string
	PLZ string
	Ort string
	Straße string

}

type DonateData struct {
	Name string
	Buchungsdatum string
	Betrag string


}

type Deposit []DonateData

type PageContent struct {
	HeaderLine string
	ICFName  string
	ICFAdress string
	ICFEmail string
	HeadlineSammelb string
	SammelbContentBeforBorder string
	SammelbContentAfterBorder string
	Vorsitzender string
	Finanzen string
	HeadlineHinweis string
	HinweisContent string

	AnlageSammelbest string

}


/* Main Funktions */

func PageContentBuilder() PageContent {
	file, _ := os.Open("./template/templatePDF.html")
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	pageContent := PageContent{}
	doc.Find(".content").Each(func(i int, s *goquery.Selection) {

		pageContent.HeaderLine = doc.Find(".headerline").Text()
		pageContent.ICFName = doc.Find(".icfAdress").Find("#icfname").Text()
		pageContent.ICFAdress = doc.Find(".icfAdress").Find("#icfadress").Text()
		pageContent.ICFEmail = doc.Find(".icfAdress").Find("#icfemail").Text()

		pageContent.HeadlineSammelb = doc.Find(".sammelbest").Find("#headline").Text()

		pageContent.SammelbContentBeforBorder = strings.TrimSpace(doc.Find(".sammelbest").Find("#contentbefor").Text())
		pageContent.SammelbContentAfterBorder = strings.TrimSpace(doc.Find(".sammelbest").Find("#contentafter").Text())

		pageContent.Vorsitzender = doc.Find(".unterschrift").Find("#vorsitzender").Text()
		pageContent.Finanzen = doc.Find(".unterschrift").Find("#finanzen").Text()

		pageContent.HeadlineHinweis = strings.TrimSpace(doc.Find(".hinweis").Find("#headline").Text())
		pageContent.HinweisContent = strings.TrimSpace(doc.Find(".hinweis").Find("#content").Text())

		pageContent.AnlageSammelbest = doc.Find(".anlage").Find("#headline").Text()
	})
	return pageContent
}

func PageBuilder(pC PageContent, ad AdressData , pD []DonateData, summe string, period string) {

	/*
	HeaderLine string
	ICFName  string
	ICFAdress string
	ICFEmail string
	HeadlineSammelb string
	SammelbContentBeforBorder string
	SammelbContentAfterBorder string
	Vorsitzender string
	Finanzen string
	HeadlineHinweis string
	HinweisContentOne string
	HinweisContentTwo string
	AnlageSammelbest string
	 */
	//adressDate := GetNameAdressDataFromCSV()

	current_data := time.Now().Local()


	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("courier", "", 6)
	umlautTranslater := pdf.UnicodeTranslatorFromDescriptor("")


	pdf.CellFormat(0,20, umlautTranslater(pC.HeaderLine), "", 1, "", false, 0, "")
	pdf.CellFormat(0,10, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0,6, umlautTranslater(pC.ICFName), "", 1, "C", false, 0, "")


	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0,6, umlautTranslater(pC.ICFAdress), "", 1, "C", false, 0, "")
	pdf.CellFormat(0,6, umlautTranslater(pC.ICFEmail), "", 1, "C", false, 0, "")
	pdf.CellFormat(0,20, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0,6, umlautTranslater(ad.Namen), "", 1, "L", false, 0, "")
	pdf.CellFormat(0,6, umlautTranslater(ad.Straße), "", 1, "L", false, 0, "")
	pdf.CellFormat(0,6, umlautTranslater(ad.Ort), "", 1, "L", false, 0, "")
	pdf.CellFormat(0,15, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0,6, umlautTranslater(pC.HeadlineSammelb), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0,6, umlautTranslater(pC.SammelbContentBeforBorder), "", "L",  false)
	pdf.CellFormat(0,5, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0,6, umlautTranslater("Name und Anschrift des Zuwendenden:"), "L,T,R", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0,6, umlautTranslater(ad.Namen + ", " + ad.Straße + ", " + ad.PLZ + ", " + ad.Ort), "L,B,R", 1, "C", false, 0, "")
	pdf.CellFormat(0,5, "", "", 1, "", false, 0, "")


	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(63.3, 6, "Gesamtbetrag der Zuwendung - in Ziffern -", "L,T,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, "- in Buchstaben -", "L,T,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater("Zeitraum der Sammelbestätigung:"), "L,T,R", 0, "C", false, 0, "")
	pdf.Ln(-1)

	//TODO: Gesamtbetrag
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(63.3, 6, summe + umlautTranslater(" €"), "L,B,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, "", "L,B,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater(period), "L,B,R", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0,8, "", "", 1, "", false, 0, "")




	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0,6, umlautTranslater(pC.SammelbContentAfterBorder), "", "L",  false)
	pdf.CellFormat(0,15, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(50, 6, umlautTranslater("Nürnberg, " + current_data.Format("02.01.2006")), "", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater(pC.Vorsitzender), "T", 0, "C", false, 0, "")
	pdf.CellFormat(2, 2, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater(pC.Finanzen), "T", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0,8, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0,6, umlautTranslater(pC.HeadlineHinweis), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0,6, umlautTranslater(pC.HinweisContent), "", "L",  false)
	pdf.CellFormat(0,5, "", "", 1, "", false, 0, "")


	pdf.AddPage()

	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(45,6, umlautTranslater(pC.AnlageSammelbest), "", 0, "", false, 0, "")
	pdf.CellFormat(45, 6, umlautTranslater("vom " + current_data.Format("02.01.2006")), "", 0, "R", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0,5, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(35, 4, umlautTranslater(ad.Namen), "", 1, "", false, 0, "")
	pdf.CellFormat(35, 4, umlautTranslater(ad.Straße), "", 1, "", false, 0, "")
	pdf.CellFormat(20, 4, umlautTranslater(ad.PLZ), "", 1, "", false, 0, "")
	pdf.CellFormat(20, 4, umlautTranslater(ad.Ort), "", 1, "", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "B", 6)
	pdf.CellFormat(45, 6, umlautTranslater("Datum der Zuwendung"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.CellFormat(45, 6, umlautTranslater("Art der Zuwendung"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.CellFormat(50, 6, umlautTranslater("Verzicht auf die Erstattung von Aufwendungen"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.CellFormat(45, 6, umlautTranslater("Betrag"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 6)
	for x := range pD {
		pdf.CellFormat(45, 6, umlautTranslater(pD[x].Buchungsdatum), "L,T,R,B", 0, "C", false, 0, "")
		pdf.CellFormat(45, 6, umlautTranslater("Geldzuwendung"), "L,B,R", 0, "C", false, 0, "")
		pdf.CellFormat(50, 6, umlautTranslater("nein"), "L,B,R", 0, "C", false, 0, "")
		pdf.CellFormat(45, 6,umlautTranslater(strings.Replace(pD[x].Betrag,".",",",1) + " €"), "L,B,R", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.CellFormat(0,8, "", "", 1, "", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(45, 6, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(45, 6, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(50, 6, umlautTranslater("Gesamtsumme"), "", 0, "C", false, 0, "")
	pdf.CellFormat(45, 6,umlautTranslater(summe + " €"), "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	err := pdf.OutputFileAndClose("./pdf/" + ad.Namen +".pdf")
	if err != nil {
		fmt.Print(err)
	}



}

/* Help Funcitons */


func GetNameAdressDataFromCSV() []AdressData {
	// Load a TXT file.
	var adressData []AdressData
	f, _ := os.Open("./csv/adressenUTF8.csv")
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {

		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if len(record) == 2 {
			recordData := "" + record[0] + "" + record[1]

			extract := strings.Split(recordData, ";")

			data := AdressData{Namen: extract[0], Firma: extract [1], PLZ: extract[2], Ort: extract[3], Straße: extract[4] }
			adressData = append(adressData, data)
		}

		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.




	}


	logrus.Infoln("Get the Adresses: DONE")
	return adressData

}

func GetDonateData() []DonateData {

	var donateData []DonateData
	f, _ := os.Open("./csv/buchungUTF8.csv")
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if len(record) == 2 {
			recordData := "" + record[0] + "" + record[1]

			extract := strings.Split(recordData, ";")

			data := DonateData{Name: extract[3], Buchungsdatum: extract [0], Betrag: extract[7] }
			donateData = append(donateData, data)

		}

	}

	logrus.Infoln("Get the Adresses: DONE")
	return donateData
}

func gesamtBetrag(donatePersonData []DonateData) string{

	summe := 0.0
	for x := range donatePersonData {
		s, _ := strconv.ParseFloat(donatePersonData[x].Betrag, 2)
		summe  += s
	}

	summeStringPoint := strconv.FormatFloat(summe,'f',-1,32)
	summeStringKomma := strings.Replace(summeStringPoint,".",",",1)
	return  summeStringKomma
}

func searchDonateContent(name string, donateData []DonateData) []DonateData{

	var spender []DonateData
	fmt.Println("INSIDE")
	for x := range donateData {

		if name == donateData[x].Name {
			spender = append(spender, DonateData{Name: donateData[x].Name, Betrag: donateData[x].Betrag, Buchungsdatum: donateData[x].Buchungsdatum})
		}
	}

	return spender
}

func CSVtoUTF8(data string) {
	f, err := os.Open("./csv/" + data + ".csv")
	if err != nil {
		// handle file open error
	}
	out, err := os.Create("./csv/" + data + "Utf8.csv")
	if err != nil {
		// handler error
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)

	io.Copy(out, r)

	out.Close()
	f.Close()
}

func Period( dep []DonateData) string {

	period := dep[0].Buchungsdatum + " - " + dep[len(dep)-1].Buchungsdatum

	return period
}

func PersonExistsDepositChecker(dep []DonateData) bool {
	if len(dep) > 0 {
		return true
	}

	return false
}

func WriteToFile(fileName string, value string) {
	f, err := os.OpenFile("./log/" + fileName + ".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(value); err != nil {
		panic(err)
	}

}

//Needed Sort Libraryfunctions for overriding!
func (a Deposit) Len() int           { return len(a) }
func (a  Deposit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a  Deposit) Less(i, j int) bool { return a[i].Buchungsdatum < a[j].Buchungsdatum }





/* MAIN */

func main() {


	//CSVtoUTF8("adressen")
	//CSVtoUTF8("buchung")
	pC := PageContentBuilder()

	adressSlice := GetNameAdressDataFromCSV()
	donateData:= GetDonateData()
	for x := range adressSlice {

		donatePersonData := searchDonateContent(adressSlice[x].Namen, donateData)
		summeString := gesamtBetrag(donatePersonData)
		sort.Sort(Deposit(donatePersonData))
		check := PersonExistsDepositChecker(donatePersonData)
		if check == true {
			period := Period(donatePersonData)
			PageBuilder(pC, adressSlice[x], donatePersonData, summeString, period)
		} else {
			WriteToFile("NoDeposit", adressSlice[x].Namen)
		}
	}



}