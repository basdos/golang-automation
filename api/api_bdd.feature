Feature: Login API

  Scenario: Login API Success
    Given I have a username "emilys" and password "emilyspass"
    When I send a POST request to endpoint "/auth/login"
    And I verify response code
    Then I verify response body