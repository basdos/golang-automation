# Golang Automation

This is project about Automation Test using Go Programming Language

# Prerequisites:

Install Go: Download and install Go from Go's official website.__

Install Go Packages:__
    - testify: For writing assertions and test cases.__
    - selenium: For web browser automation.__
    - appium Go bindings (go-selenium can be used with Appium).__

npm install -g appium # for mobile test automation__

# Automation Api 

https://dummyjson.com with login endpoint__
Use Gherkin BDD cucumber, file name contains the word "bdd"

Command to run  "go test -run ^TestApiLogin$ ./api"

without Gherkin BDD cucumber , file name not contains the word "bdd"

Command to run "go test ./api/api_test.go"
