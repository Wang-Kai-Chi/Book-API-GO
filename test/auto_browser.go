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
	au.Assert(err)
	au.Playwright = pw

	browser, err := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(au.IsHideBrowser),
		},
	)
	au.Assert(err)
	au.Browser = browser

	return au
}

func (au AutoBrowser) Assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (au AutoBrowser) Start() playwright.Page {
	page, err := au.Browser.NewPage()
	au.Assert(err)

	_, err = page.Goto(au.Url)
	au.Assert(err)
	return page
}

func (au AutoBrowser) End() {
	au.Assert(au.Browser.Close())
	au.Assert(au.Playwright.Stop())
}
