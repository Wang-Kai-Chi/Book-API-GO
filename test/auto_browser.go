package test

import (
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
	Assert(err)
	au.Playwright = pw

	browser, err := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(au.IsHideBrowser),
		},
	)
	Assert(err)
	au.Browser = browser

	return au
}

func (au AutoBrowser) Start() playwright.Page {
	page, err := au.Browser.NewPage()
	Assert(err)

	_, err = page.Goto(au.Url)
	Assert(err)
	return page
}

func (au AutoBrowser) End() {
	Assert(au.Browser.Close())
	Assert(au.Playwright.Stop())
}
