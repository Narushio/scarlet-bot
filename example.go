package main

//func main() {
//	go setHttpServer(8080)
//	chromium := browser.LaunchChromium()
//	page, err := chromium.NewPage()
//	if err != nil {
//		log.Fatalf("create page: %v", err)
//	}
//	if _, err = page.Goto("http://localhost:8080/qqx5/duo-burst-table.html"); err != nil {
//		log.Fatalf("goto: %v", err)
//	}
//
//	jsonData, err := json.Marshal(model.DuoBurstTableList[0])
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	_, err = page.Evaluate(fmt.Sprintf("initDom(%s)", string(jsonData)))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	time.Sleep(500 * time.Millisecond)
//	_, err = page.Locator("div.main").Screenshot(playwright.LocatorScreenshotOptions{
//		Path: playwright.String("tmp/example.png"),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//}
