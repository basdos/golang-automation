package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

var (
	baseURL      = "https://dummyjson.com" // json dummy base url
	username     string
	password     string
	response     *http.Response
	responseBody map[string]interface{}
)

// Step Definitions
func iHaveAUsernameAndPassword(user, pass string) error {
	username = user
	password = pass
	return nil
}

func iSendAPostRequestToEndpoint(endpoint string) error {
	url := baseURL + endpoint

	// Create the request body
	reqBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return err
	}
	// Send the POST request
	response, err = http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	// Read and parse the response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	jsonString := string(bodyBytes)
	fmt.Println("Response Body : " + jsonString)

	json.Unmarshal(bodyBytes, &responseBody)
	return nil
}

func iVerifyResponseCode() error {
	assert.Equal(nil, 200, response.StatusCode)
	return nil
}

func iVerifyResponseBody() error {
	assert.Equal(nil, 1, int(responseBody["id"].(float64)))
	assert.Equal(nil, "emilys", responseBody["username"].(string))
	fmt.Println("Access Token : " + responseBody["accessToken"].(string))
	return nil
}

// Initialize the godog context
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have a username "([^"]*)" and password "([^"]*)"$`, iHaveAUsernameAndPassword)
	ctx.Step(`^I send a POST request to endpoint "([^"]*)"$`, iSendAPostRequestToEndpoint)
	ctx.Step(`^I verify response code`, iVerifyResponseCode)
	ctx.Step(`^I verify response body`, iVerifyResponseBody)
}
