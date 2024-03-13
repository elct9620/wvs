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

  Scenario: The joined match will be returned
    Given the session id is "d7ae7356-2d91-47f7-81bd-428c40bf55c3"
    And there have some match
      """
        [
          {
            "id": "6b590342-b50d-4f5a-a2d3-f6899014a37b",
            "players": [
              {
                "team": "slime",
                "id": "3b51562c-62dc-4fba-b4d2-87797a186d1d"
              }
            ]
          },
          {
            "id": "83b03486-d1ed-44cf-b3b8-694ea73de901",
            "players": [
              {
                "team": "walrus",
                "id": "d7ae7356-2d91-47f7-81bd-428c40bf55c3"
              }
            ]
          }
        ]
      """
    When I make a POST request to "/api/match" with body
      """
      {
        "team": "slime"
      }
      """
    Then the response JSON should has "match_id" with value "83b03486-d1ed-44cf-b3b8-694ea73de901"
    And the response status code should be 200

  Scenario: I can join a existing match
    Given there have some match
      """
        [
          {
            "id": "83b03486-d1ed-44cf-b3b8-694ea73de901",
            "players": [
              {
                "team": "walrus",
                "id": "61a607f3-3373-4a21-85f9-8d94f8738acf"
              }
            ]
          }
        ]
      """
    When I make a POST request to "/api/match" with body
      """
      {
        "team": "slime"
      }
      """
    Then the response JSON should has "match_id" with value "83b03486-d1ed-44cf-b3b8-694ea73de901"
    And the response status code should be 200
