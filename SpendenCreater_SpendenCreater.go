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

func PageBuilder(pC PageContent, ad AdressData) {

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
	pdf.CellFormat(63.3, 6, "TEST", "L,B,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, "TEST", "L,B,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, "TEST", "L,B,R", 0, "C", false, 0, "")
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


	err := pdf.OutputFileAndClose("./pdf/" + ad.Namen +".pdf")
	if err != nil {
		fmt.Print(err)
	}



}


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
	// Load a TXT file.
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

func searchDonateContent(name string, donateData []DonateData) []DonateData{

	var spender []DonateData
	fmt.Println("INSIDE")
	for x := range donateData {
		fmt.Println(donateData[x].Name)
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




func main() {



	/*
	CSVtoUTF8("adressen")
	CSVtoUTF8("buchung")
	pC := PageContentBuilder()

	adressSlice := GetNameAdressDataFromCSV()
	for x := range adressSlice {
		PageBuilder(pC, adressSlice[x])
	}
	*/
	fmt.Println("HLLO")
	data := GetDonateData()
	slicey := searchDonateContent("Sturm Robert", data)
	fmt.Println(slicey)

}