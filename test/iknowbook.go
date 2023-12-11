package test

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

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

func (app Iknowbook) MustFindUserMenu(page playwright.Page) bool {
	menu := page.Locator("#dropdownUser")
	bo, err := menu.IsEnabled()
	Assert(err)
	return bo
}

func (app Iknowbook) MustOpenUserMenu(page playwright.Page) bool {
	menu := page.Locator("#dropdownUser")
	Assert(menu.Click())
	menuopen, err := page.Locator("#userDroplist").IsEnabled()
	Assert(err)
	return menuopen
}

func (app Iknowbook) Login(page playwright.Page) {
	menuopen := app.MustOpenUserMenu(page)
	if menuopen {
		Assert(page.Locator("#userDroplist>ul>li:nth-child(1)").Click())

		Assert(page.Locator("#email").Fill("testA@mail.com"))
		Assert(page.Locator("#password").Fill("testpassword"))
		Assert(page.Locator("#submit").Click())
	} else {
		log.Fatal("Can't open user menu.")
	}
}

func (app Iknowbook) Logout(page playwright.Page) {
	au := app.AutoBrowser
	Assert(page.Locator("#dropdownUser").Click())
	Assert(page.Locator("#userDroplist>ul>li:nth-child(3)").Click())
	au.End()
}
