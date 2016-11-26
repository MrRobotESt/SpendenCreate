package helper

import (
	"encoding/csv"
	"bufio"
	"io"
	"strings"
	"github.com/Sirupsen/logrus"
	st "../structs"
	"os"
)

func GetDonateData() []st.DonateData {

	var donateData []st.DonateData

	f, _ := os.Open("./csv/buchungUTF8.csv")
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if len(record) == 2 {
			recordData := "" + record[0] + "" + record[1]

			extract := strings.Split(recordData, ";")

			data := st.DonateData{Namen: extract[3], Buchungsdatum: extract [0], Betrag: extract[7] }
			donateData = append(donateData, data)

		}

	}

	logrus.Infoln("Get the DonateData: DONE!")
	return donateData
}