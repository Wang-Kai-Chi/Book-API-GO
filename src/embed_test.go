package main

import (
	"embed"
	"testing"

	"iknowbook.com/handler"
)

//go:embed test.txt
var fs embed.FS

func TestEmbed(t *testing.T) {
	data, err := fs.ReadFile("test.txt")
	if err == nil {
		println(string(data))
	} else {
		t.Fatal(err)
	}
}

func TestReadTxt(t *testing.T) {
	data := handler.ReadFileAsString("test.txt")
	t.Log(data)
}
