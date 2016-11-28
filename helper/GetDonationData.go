package helper

import (
	st "../structs"
	"bufio"
	"encoding/csv"
	"github.com/Sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetDonateData() []st.DonateData {

	var donateData []st.DonateData
	var recordData string

	f, _ := os.Open("./csv/buchungUTF8.csv")
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		//Sometimes there are Problems with "," which increas the record slice from 2 to 7 or 9
		// That is the reason why i have to check different cases!
		// z.B:  04.01.2016;_HVB Spendenkonto;2496;Sturm, Robert;Spende von XYZ, Strasse 24, 90443 Nuernberg;Spenden:05 Spende 10ter;Spendeneinnahmen;75
		//Problem: Spend von XYZ, => cut to a new Sliece ! INCREASE THE AMOUNT!
		if len(record) >= 2 {

			for i := 0; i < len(record)-2; i++ {
				recordData += "" + record[i]
			}

			a, err := strconv.Atoi(record[len(record)-2])
			if err != nil {
				recordData += "" + record[len(record)-2] + "" + record[len(record)-1]
			} else {
				recordData += "" + record[len(record)-2] + "." + record[len(record)-1]
				logrus.Infoln(a)
			}

			extract := strings.Split(recordData, ";")
			data := st.DonateData{Namen: extract[3], Buchungsdatum: extract[0], Betrag: extract[len(extract)-1]}

			donateData = append(donateData, data)

		}

	}

	logrus.Infoln("Get the DonateData: DONE!")
	return donateData
}
