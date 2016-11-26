package spendenCreator


import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
	"github.com/jung-kurt/gofpdf"
	"fmt"
	st "../structs"
	"os"
)


/* Main Funktions */


//Fill the PageContent struct with the htmlTemplate Data
func PageContentBuilder() st.PageContent {
	file, _ := os.Open("./template/templatePDF.html")
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	pageContent := st.PageContent{}
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


//Create the PDF from PageContent Data
func PageBuilder(pC st.PageContent, ad st.AdressData , pD []st.DonateData, summe string, period string) {

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


	current_data := time.Now().Local()


	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("courier", "", 6)
	umlautTranslater := pdf.UnicodeTranslatorFromDescriptor("")

	//Header Line
	pdf.CellFormat(0,20, umlautTranslater(pC.HeaderLine), "", 1, "", false, 0, "")
	pdf.CellFormat(0,10, "", "", 1, "", false, 0, "")


	// ICF Adress
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0,6, umlautTranslater(pC.ICFName), "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0,6, umlautTranslater(pC.ICFAdress), "", 1, "C", false, 0, "")
	pdf.CellFormat(0,6, umlautTranslater(pC.ICFEmail), "", 1, "C", false, 0, "")
	pdf.CellFormat(0,20, "", "", 1, "", false, 0, "")

	// Donater Adress Data
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0,6, umlautTranslater(ad.Namen), "", 1, "L", false, 0, "")
	pdf.CellFormat(0,6, umlautTranslater(ad.Straße), "", 1, "L", false, 0, "")
	pdf.CellFormat(0,6, umlautTranslater(ad.Ort), "", 1, "L", false, 0, "")
	pdf.CellFormat(0,15, "", "", 1, "", false, 0, "")

	//Sammelbestätigung ContentBefor dynamic table content
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

	// Table Summ of Donation and TimePeriod
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(63.3, 6, "Gesamtbetrag der Zuwendung - in Ziffern -", "L,T,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, "- in Buchstaben -", "L,T,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater("Zeitraum der Sammelbestätigung:"), "L,T,R", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(63.3, 6, summe + umlautTranslater(" €"), "L,B,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, "", "L,B,R", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater(period), "L,B,R", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0,8, "", "", 1, "", false, 0, "")

	//Sammelbestätigung ContentAfter dynamic table content
	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0,6, umlautTranslater(pC.SammelbContentAfterBorder), "", "L",  false)
	pdf.CellFormat(0,15, "", "", 1, "", false, 0, "")

	//Signature Vorsitzender, Finanzen and Date
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(50, 6, umlautTranslater("Nürnberg, " + current_data.Format("02.01.2006")), "", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater(pC.Vorsitzender), "T", 0, "C", false, 0, "")
	pdf.CellFormat(2, 2, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(63.3, 6, umlautTranslater(pC.Finanzen), "T", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0,8, "", "", 1, "", false, 0, "")


	//Hinweis (Footer)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0,6, umlautTranslater(pC.HeadlineHinweis), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.MultiCell(0,6, umlautTranslater(pC.HinweisContent), "", "L",  false)
	pdf.CellFormat(0,5, "", "", 1, "", false, 0, "")


	/* NewPage for TableContent */
	pdf.AddPage()


	//Headline and Date
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(45,6, umlautTranslater(pC.AnlageSammelbest), "", 0, "", false, 0, "")
	pdf.CellFormat(45, 6, umlautTranslater("vom " + current_data.Format("02.01.2006")), "", 0, "R", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(0,5, "", "", 1, "", false, 0, "")

	//Donater Data
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(35, 4, umlautTranslater(ad.Namen), "", 1, "", false, 0, "")
	pdf.CellFormat(35, 4, umlautTranslater(ad.Straße), "", 1, "", false, 0, "")
	pdf.CellFormat(20, 4, umlautTranslater(ad.PLZ), "", 1, "", false, 0, "")
	pdf.CellFormat(20, 4, umlautTranslater(ad.Ort), "", 1, "", false, 0, "")
	pdf.Ln(-1)

	//Static Table Headline
	pdf.SetFont("Arial", "B", 6)
	pdf.CellFormat(45, 6, umlautTranslater("Datum der Zuwendung"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.CellFormat(45, 6, umlautTranslater("Art der Zuwendung"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.CellFormat(50, 6, umlautTranslater("Verzicht auf die Erstattung von Aufwendungen"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.CellFormat(45, 6, umlautTranslater("Betrag"), "L,T,R,B", 0, "C", false, 0, "")
	pdf.Ln(-1)

	//Dynamic Donater Period donation Data!
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

	//Create the PDF
	err := pdf.OutputFileAndClose("./pdf/" + ad.Namen +".pdf")
	if err != nil {
		fmt.Print(err)
	}



}









