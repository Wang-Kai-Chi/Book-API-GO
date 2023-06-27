package handler

import (
	"os"
	"testing"
)

func GetFileName() string {
	return "../output/iknowbook.txt"
}

func GetContent() string {
	return "Welcome to IKnowBook Store~!"
}

func TestCreateFile(t *testing.T) {
	fileName := GetFileName()
	CreateFile(fileName, GetContent())

	_, err := os.Stat(fileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFileAsString(t *testing.T) {
	result := ReadFileAsString(GetFileName())

	if result == "" {
		t.Fatal()
	}
}
