package test

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

type AutoBrowser struct {
	Playwright    *playwright.Playwright
	Browser       playwright.Browser
	Url           string
	IsHideBrowser bool
}

func NewAutoBrowser(url string, isHideBrowser bool) AutoBrowser {
	au := AutoBrowser{
		Url:           url,
		IsHideBrowser: isHideBrowser,
	}
	pw, err := playwright.Run()
	fatal(err)
	au.Playwright = pw

	browser, err := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(au.IsHideBrowser),
		},
	)
	fatal(err)
	au.Browser = browser

	return au
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (au AutoBrowser) Start() playwright.Page {
	page, err := au.Browser.NewPage()
	fatal(err)

	_, err = page.Goto(au.Url)
	fatal(err)
	return page
}

func (au AutoBrowser) End() {
	fatal(au.Browser.Close())
	fatal(au.Playwright.Stop())
}
