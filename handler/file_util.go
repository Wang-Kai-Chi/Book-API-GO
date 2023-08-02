package handler

import (
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
	return string(MustReadFile(fileName)[:])
}

func MustReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	return data
}
