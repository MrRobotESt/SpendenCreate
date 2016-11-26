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

//Get the AdressData and put them into the AdressData struct

func GetNameAdressDataFromCSV() []st.AdressData {

	var adressData []st.AdressData
	f, _ := os.Open("./csv/adressenUTF8.csv")

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

			data := st.AdressData{Namen: extract[0], Firma: extract [1], PLZ: extract[2], Ort: extract[3], Straße: extract[4] }
			adressData = append(adressData, data)
		}




	}


	logrus.Infoln("Get the Adresses: DONE!")
	return adressData

}