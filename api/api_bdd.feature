Feature: Login API

  Scenario: Login Success
    Given I have a username "emilys" and password "emilyspass"
    When I send a POST request to endpoint "/auth/login"
    Then I verify response code
    And I verify response body