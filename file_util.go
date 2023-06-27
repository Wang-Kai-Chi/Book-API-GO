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
	return string(ReadFile(fileName)[:])
}

func ReadFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	return data
}
