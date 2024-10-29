# Golang Automation

This is project about Automation Test using Go Programming Language

# Prerequisites:

Install Go: Download and install Go from Go's official website.<br />

Install Go Packages:<br />
    - testify: For writing assertions and test cases.<br />
    - selenium: For web browser automation.<br />
    - appium Go bindings (go-selenium can be used with Appium).<br />

npm install -g appium # for mobile test automation <br />

# Automation Api 

https://dummyjson.com with login endpoint <br />
Use Gherkin BDD cucumber, file name contains the word "bdd"

Command to run  "go test -run ^TestApiLogin$ ./api"

without Gherkin BDD cucumber , file name not contains the word "bdd"

Command to run "go test ./api/api_test.go"
