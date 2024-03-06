Feature: Me API
  Scenario: When player get UUID than return UUID
    Given the session id is "d7ae7356-2d91-47f7-81bd-428c40bf55c3"
    When I make a GET request to "/api/me"
    Then the response body should be a valid JSON
      """
      {
        "id": "d7ae7356-2d91-47f7-81bd-428c40bf55c3"
      }
      """
    And the response status code should be 200
