Feature: Main Scene
  Scenario: The main scene is served
    When I make a GET request to "/"
    Then the response status code should be 200
