package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"

	"encoding/csv"

	"golang.org/x/text/encoding/charmap"
)

type AdressData struct {
	Namen string
	Firma string
	PLZ string
	Ort string
	Straße string

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
	HinweisContentOne string
	HinweisContentTwo string
	AnlageSammelbest string

}



func PageContentBuilder() PageContent {
	pageContent := PageContent{}
	pageContent.HeaderLine = "Aussteller (Bezeichnung und Anschrift der steuerbegünstigten Einrichtung)"
	pageContent.ICFName = "ICF Nürnberg e.V."
	pageContent.ICFAdress = "Königstraße 57; 90402 Nürnberg"
	pageContent.ICFEmail = "E-Mail: buchhaltung@icfn.de"
	pageContent.HeadlineSammelb = "Sammelbestätigung über Geldzuwendungen/Mitgliedsbeiträge"

	pageContent.SammelbContentBeforBorder = "im Sinne des § 10b des Einkommensteuergesetzes an eine der in § 5 Abs. 1 Nr. 9 des Körperschaftsteuergesetzes bezeichneten"
	pageContent.SammelbContentBeforBorder += "Körperschaften, Personenvereinigungen oder Vermögensmassen"

	pageContent.SammelbContentAfterBorder = "Wir sind wegen Förderung mildtätiger Zwecke sowie folgender gemeinnütziger Zwecke: Förderung der Religion, der Jugendhilfe."
	pageContent.SammelbContentAfterBorder += "Es wird bestätigt, dass die Zuwendungen nur zur Förderung mildtätiger Zwecke und folgender gemeinnütziger Zwecke: Förderung der"
	pageContent.SammelbContentAfterBorder += "Religion, der Jugendhilfe, der internationalen Gesinnung, der Toleranz auf allen Gebieten der Kultur und des"
	pageContent.SammelbContentAfterBorder += "Völkerverständigungsgedankens sowie des Schutzes von Ehe und Familie verwendet werden."
	pageContent.SammelbContentAfterBorder += "Es wird bestätigt, dass es sich nicht um einen Mitgliedsbeitrag handelt, dessen Abzug nach § 10b Abs. 1 des Einkommensteuergesetzes ausgeschlossen ist."
	pageContent.SammelbContentAfterBorder += "Es wird bestätigt, dass über die in der Gesamtsumme enthaltenen Zuwendungen keine weiteren Bestätigungen, weder formelle"
	pageContent.SammelbContentAfterBorder += "Zuwendungsbestätigungen noch Beitragsquittungen oder Ähnliches ausgestellt wurden und werden."
	pageContent.SammelbContentAfterBorder += "Ob es sich um den Verzicht auf Erstattung von Aufwendungen handelt, ist der Anlage zur Sammelbestätigung zu entnehmen."

	pageContent.Vorsitzender = "Vorsitzender: Daniel Kalupner"
	pageContent.Finanzen     = "Finanzen: Steffen Wentzel"
	pageContent.HeadlineHinweis = "Hinweis:"

	pageContent.HinweisContentOne = "Wer vorsätzlich oder grob fahrlässig eine unrichtige Zuwendungsbestätigung erstellt oder veranlasst, dass Zuwendungen nicht zu den in der"
	pageContent.HinweisContentOne += "Zuwendungsbestätigung angegebenen steuerbegünstigten Zwecken verwendet werden, haftet für die entgangene Steuer (§ 10b Abs. 4 EStG, § 9 Abs. 3 KStG, § 9 Nr. 5 GewStG)."

	pageContent.AnlageSammelbest = "Anlage zur Sammelbestätigung"

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
	pdf.CellFormat(0,20, "", "", 1, "", false, 0, "")

	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0,6, umlautTranslater(pC.HeadlineSammelb), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 12)
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


	err := pdf.OutputFileAndClose(ad.Namen +".pdf")
	if err != nil {
		fmt.Print(err)
	}

}


func GetNameAdressDataFromCSV() []AdressData {
	// Load a TXT file.
	var adressData []AdressData
	f, _ := os.Open("adressenUTF8.csv")
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

	fmt.Println(adressData[0].Ort)
	fmt.Println("ÜÜÜÜÜÜ")
	return adressData

}


func GetDonateData() {

}

func CSVtoUTF8() {
	f, err := os.Open("adressen.csv")
	if err != nil {
		// handle file open error
	}
	out, err := os.Create("adressenUtf8.csv")
	if err != nil {
		// handler error
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)

	io.Copy(out, r)

	out.Close()
	f.Close()
}

func main() {
	pC := PageContentBuilder()

	adressSlice := GetNameAdressDataFromCSV()
	for x := range adressSlice {
		PageBuilder(pC, adressSlice[x])
	}

	//CSVtoUTF8()

}