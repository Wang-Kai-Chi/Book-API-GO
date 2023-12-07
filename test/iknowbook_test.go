package test

import "testing"

func TestLoginLogout(t *testing.T) {
	var au AutoBrowser
	au = NewAutoBrowser("http://localhost:8081", false)
	page := au.Start()
	au.Assert(page.Locator("#dropdownUser").Click())
	au.Assert(page.Locator("#userDroplist>ul>li:nth-child(1)").Click())
	au.Assert(page.Locator("#email").Fill("testA@mail.com"))
	au.Assert(page.Locator("#password").Fill("testpassword"))
	au.Assert(page.Locator("#submit").Click())
	au.Assert(page.Locator("#dropdownUser").Click())
	au.Assert(page.Locator("#userDroplist>ul>li:nth-child(3)").Click())
	au.End()
}
