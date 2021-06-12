package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
func GetFromStdin(fromStdin string) string {
	var (
		value string = ""
	)
	if fromStdin != "" {
		value = fromStdin
	}
	return value
}

func GetFromFile(fromFile string) string {
	var (
		returnValue string = ""
	)
	if fromFile != "" {
		var f, err = ioutil.ReadFile(fromFile)
		Check(err)
		returnValue = string(f)
	}
	return returnValue
}

func WriteToStdout(toStdout bool, result string) bool {
	if toStdout == true {
		var (
			err error
		)
		_, err = fmt.Fprintln(os.Stdout, result)
		Check(err)
		return true
	}
	return false
}
func WriteToFile(toFile string, value string) {
	if toFile != "" {
		var (
			mydata []byte
		)
		mydata = []byte(value)
		err := ioutil.WriteFile(toFile, mydata, 0777)
		Check(err)
	}
}
