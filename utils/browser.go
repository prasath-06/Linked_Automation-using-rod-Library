package utils

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// To start Chrome browser
func StartBrowser() (*rod.Browser, *rod.Page) {
	url := launcher.New().
		//visible browser
		Headless(false).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage()

	return browser, page
}
