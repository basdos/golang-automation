package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/usr/local/bin/chromedriver" // Path to ChromeDriver
	port             = 8080                          // Port for WebDriver
)

func TestWeb(t *testing.T) {
	// Start a new WebDriver service
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver service: %v", err)
	}
	defer service.Stop()

	// Set the WebDriver capabilities (e.g., for Chrome)
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// Connect to the WebDriver instance running locally
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error connecting to the WebDriver: %v", err)
	}
	defer wd.Quit()

	// Open the SauceDemo login page
	if err := wd.Get("https://www.saucedemo.com/"); err != nil {
		log.Fatalf("Error opening SauceDemo website: %v", err)
	}

	// Find the username field and enter a username
	usernameField, err := wd.FindElement(selenium.ByID, "user-name")
	if err != nil {
		log.Fatalf("Error finding username field: %v", err)
	}
	usernameField.SendKeys("standard_user")

	// Find the password field and enter a password
	passwordField, err := wd.FindElement(selenium.ByID, "password")
	if err != nil {
		log.Fatalf("Error finding password field: %v", err)
	}
	passwordField.SendKeys("secret_sauce")

	// Find the login button and click it
	loginButton, err := wd.FindElement(selenium.ByID, "login-button")
	if err != nil {
		log.Fatalf("Error finding login button: %v", err)
	}
	loginButton.Click()

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

	defer wd.Quit()
}
