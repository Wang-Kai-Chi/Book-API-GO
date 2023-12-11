package test

import (
	"fmt"
	"log"
	"testing"
)

func TestStart(t *testing.T) {
	var auto AutoBrowser
	auto = NewAutoBrowser("http://localhost:8081", false)
	auto.Start()
	auto.End()
}

func TestPlayWright(t *testing.T) {
	var auto AutoBrowser
	auto = NewAutoBrowser("https://news.ycombinator.com", false)
	page := auto.Start()
	entries, err := page.Locator(".athing").All()
	Assert(err)

	for i, entry := range entries {
		title, err := entry.Locator("td.title > span > a").TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	auto.End()
}
