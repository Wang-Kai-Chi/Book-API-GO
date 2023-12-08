package test

import "github.com/playwright-community/playwright-go"

type Iknowbook struct {
	AutoBrowser
}

func NewIknowBook() Iknowbook {
	return Iknowbook{
		AutoBrowser: NewAutoBrowser("http://localhost:8081", false),
	}
}

func (app Iknowbook) Start() playwright.Page {
	page := app.AutoBrowser.Start()
	return page
}

func (app Iknowbook) Login(page playwright.Page) {
	au := app.AutoBrowser
	au.Assert(page.Locator("#dropdownUser").Click())
	au.Assert(page.Locator("#userDroplist>ul>li:nth-child(1)").Click())

	au.Assert(page.Locator("#email").Fill("testA@mail.com"))
	au.Assert(page.Locator("#password").Fill("testpassword"))
	au.Assert(page.Locator("#submit").Click())
}

func (app Iknowbook) Logout(page playwright.Page) {
	au := app.AutoBrowser
	au.Assert(page.Locator("#dropdownUser").Click())
	au.Assert(page.Locator("#userDroplist>ul>li:nth-child(3)").Click())
	au.End()
}
