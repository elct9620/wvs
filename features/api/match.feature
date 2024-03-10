Feature: Match
  Scenario: I can create a new match
    When I make a POST request to "/api/match" with body
      """
      {
        "team": "slime"
      }
      """
    Then the response body should be a valid JSON
      """
      {
        "ok": true
      }
      """
    And the response status code should be 200
