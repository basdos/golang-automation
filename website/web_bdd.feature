Feature: Login Web

  Scenario: Login Web Success
    Given I open browser
    When I go to login page
    And I enter valid credentials username "standard_user" and password "secret_sauce"
    Then I validate success login