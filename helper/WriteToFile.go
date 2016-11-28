package helper

import (
	"os"
)

func WriteToFile(fileName string, value string) {
	f, err := os.OpenFile("./log/"+fileName+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(value); err != nil {
		panic(err)
	}

	defer f.Close()

}
