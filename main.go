package main

import (
	"fmt"
	"time"

	"Linkedin_Automation/utils"
)

func main() {

	fmt.Println("Starting LinkedIn Automation ")

	// Start Browser
	browser, page := utils.StartBrowser()
	defer browser.MustClose()

	// Open LinkedIn Login Page
	fmt.Println("Loading LinkedIn Login Page")
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()
	utils.HumanDelay(3, 5)

	// Enter Login Credentials
	fmt.Println("Entering credentials")
	page.MustElement("#username").MustInput(utils.Credentials.Username)
	utils.HumanDelay(1, 2)

	page.MustElement("#password").MustInput(utils.Credentials.Password)
	utils.HumanDelay(1, 2)

	page.MustElement("button[type=submit]").MustClick()
	utils.HumanDelay(5, 7)

	// Verify Login
	fmt.Println("Login successful loading feed")
	page.MustWaitLoad()
	utils.HumanDelay(3, 5)

	// Go to Search Page
	fmt.Println("Navigating to search page")
	page.MustNavigate("https://www.linkedin.com/search/results/people/?keywords=software%20engineer")
	page.MustWaitLoad()
	utils.HumanDelay(3, 5)

	// Scroll Like a Human (FIXED)
	fmt.Println("Scrolling results")
	for i := 0; i < 5; i++ {
		page.MustEval(`() => window.scrollBy(0, 600)`)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Automation Completed Successfully")
	utils.HumanDelay(5, 7)
}
