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

  Scenario: I can join a existing match
    Given there have a match
      """
        {
          "id": "83b03486-d1ed-44cf-b3b8-694ea73de901",
          "players": [
            {
              "team": "walrus",
              "id": "61a607f3-3373-4a21-85f9-8d94f8738acf"
            }
          ]
        }
      """
    When I make a POST request to "/api/match" with body
      """
      {
        "team": "slime"
      }
      """
    Then the response JSON should has "match_id" with value "83b03486-d1ed-44cf-b3b8-694ea73de901"
    And the response status code should be 200
