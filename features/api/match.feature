Feature: Match
  Scenario: I can create a new match
    When I make a POST request to "/api/match" with body
      """
      {
        "team": "slime"
      }
      """
    Then the response JSON should has "match_id"
    And the response status code should be 200
