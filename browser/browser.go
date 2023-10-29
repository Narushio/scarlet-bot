package browser

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

var Chromiun playwright.Browser

func LaunchChromium() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("launch browser: %v", err)
	}
	Chromiun = browser
}
