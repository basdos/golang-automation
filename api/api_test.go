package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	baseUrl = "https://dummyjson.com" // json dummy base url
)

// Struct to represent the response object
type Login struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Id          int    `json:"id"`
	Gender      string `json:"gender"`
	AccessToken string `json:"accessToken"`
}

// Test the POST /login API
func TestApi(t *testing.T) {
	// Create a new post object to send in the request
	newPost := Login{
		Username: "emilys",
		Password: "emilyspass",
	}

	postBody, err := json.Marshal(newPost)
	if err != nil {
		t.Fatalf("Failed to request body: %v", err)
	}

	req, err := http.NewRequest("POST", baseUrl+"/auth/login", bytes.NewBuffer(postBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	assert.Equal(t, 200, resp.StatusCode, "Expected status code 200")

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	responseBody := string(body)
	// Unmarshal response into the Post struct
	var createdPost Login
	err = json.Unmarshal(body, &createdPost)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Validate the created post
	assert.Equal(t, 1, createdPost.Id, "ID verified")
	assert.Equal(t, "female", createdPost.Gender, "Gender verified")
	fmt.Println(responseBody)
	fmt.Println("Access Token : " + createdPost.AccessToken)
}
