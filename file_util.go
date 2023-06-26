package main

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
}
func ReadFileAsString(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	//return the data converted from []byte to string
	return string(data[:])
}
