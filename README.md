# Golang Automation

This is project about Automation Test using Go Programming Language

# Prerequisites:

Install Go: Download and install Go from Go's official website.
Install Go Packages:
    - testify: For writing assertions and test cases.
    - selenium: For web browser automation.
    - appium Go bindings (go-selenium can be used with Appium).
npm install -g appium # for mobile test automation

# Automation Api 

https://dummyjson.com with login endpoint
Use Gherkin BDD cucumber, file name contains the word "bdd"

Command to run  "go test -run ^TestApiLogin$ ./api"

without Gherkin BDD cucumber , file name not contains the word "bdd"

Command to run "go test ./api/api_test.go"
