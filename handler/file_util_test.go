package handler

import (
	"os"
	"testing"
)

const FILE_NAME = "../output/iknowbook.txt"
const CONTENT = "Welcome to IKnowBook Store~!"

func TestCreateFile(t *testing.T) {
	fileName := FILE_NAME
	CreateFile(fileName, CONTENT)

	_, err := os.Stat(fileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFileAsString(t *testing.T) {
	result := ReadFileAsString(FILE_NAME)

	if result == "" {
		t.Fatal()
	}
}

func TestReadFile(t *testing.T) {
	result := MustReadFile(FILE_NAME)

	if len(result) == 0 {
		t.Fatal()
	}
}
