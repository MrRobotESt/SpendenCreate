package helper

import (
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
)

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