package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cucumber/godog"
	"github.com/tebeka/selenium"
)

const (
	chromeDriver = "/usr/local/bin/chromedriver" // Path to ChromeDriver
	seleniumPort = 8080                          // Port for WebDriver
)

var wd selenium.WebDriver
var service *selenium.Service

func iOpenBrowser() error {
	// Start a new WebDriver service
	var err error
	service, err = selenium.NewChromeDriverService(chromeDriver, seleniumPort)
	if err != nil {
		log.Fatalf("error starting ChromeDriver service: %v", err)
	}

	// Set the WebDriver capabilities (e.g., for Chrome)
	caps := selenium.Capabilities{"browserName": "chrome"}

	// Connect to the WebDriver instance running locally
	wd, err = selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", seleniumPort))

	if err != nil {
		log.Fatalf("error creating WebDriver session: %v", err)
	}
	return nil
}

func iGoToLoginPage() error {
	// Open the SauceDemo login page
	if err := wd.Get("https://www.saucedemo.com/"); err != nil {
		log.Fatalf("Error opening SauceDemo website: %v", err)
	}
	return nil
}

func iEnterValidCredentials(user, pass string) error {
	// Find the username field and enter a username
	usernameField, err := wd.FindElement(selenium.ByID, "user-name")
	if err != nil {
		log.Fatalf("Error finding username field: %v", err)
	}
	usernameField.SendKeys(user)

	// Find the password field and enter a password
	passwordField, err := wd.FindElement(selenium.ByID, "password")
	if err != nil {
		log.Fatalf("Error finding password field: %v", err)
	}
	passwordField.SendKeys(pass)

	// Find the login button and click it
	loginButton, err := wd.FindElement(selenium.ByID, "login-button")
	if err != nil {
		log.Fatalf("Error finding login button: %v", err)
	}
	loginButton.Click()
	return nil
}

func iValidateSuccessLogin() error {
	// Validate login success by checking for an element on the dashboard page
	if err := wd.WaitWithTimeoutAndInterval(selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
		inventoryContainer, err := wd.FindElement(selenium.ByID, "inventory_container")
		if err != nil {
			return false, nil // Keep waiting if the element isn't found
		}
		return inventoryContainer.IsDisplayed()
	}), 10*time.Second, 500*time.Millisecond); err != nil {
		log.Fatalf("Inventory page not displayed: %v", err)
	}

	// Check for an item on the inventory page to validate successful login.
	item, err := wd.FindElement(selenium.ByID, "item_4_title_link")
	if err != nil {
		log.Fatalf("Inventory item not found: %v", err)
	}

	itemText, err := item.Text()
	if err != nil {
		log.Fatalf("Failed to get item text: %v", err)
	}

	// Verify if the login was successful
	if itemText == "Sauce Labs Backpack" {
		fmt.Println("Test Passed: Found Sauce Labs Backpack in inventory")
	} else {
		log.Fatalf("Test Failed: Expected item not found in inventory")
	}

	if wd != nil {
		wd.Quit()
	}
	if service != nil {
		service.Stop()
	}

	return nil
}

// Initialize the godog context
func InitializeScenarioWeb(ctx *godog.ScenarioContext) {
	ctx.Step(`^I open browser`, iOpenBrowser)
	ctx.Step(`^I go to login page`, iGoToLoginPage)
	ctx.Step(`^I enter valid credentials username "([^"]*)" and password "([^"]*)"$`, iEnterValidCredentials)
	ctx.Step(`^I validate success login`, iValidateSuccessLogin)
}
