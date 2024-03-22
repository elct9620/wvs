Feature: Start Battle
  Scenario: When match is full, start battle
    Given the session id is "d7ae7356-2d91-47f7-81bd-428c40bf55c3"
    And connect to the websocket
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
          }
        ]
      """
    When I make a POST request to "/api/match" with body
      """
      {
        "team": "walrus"
      }
      """
    Then the response JSON should has "match_id" with value "6b590342-b50d-4f5a-a2d3-f6899014a37b"
    And the response status code should be 200
    And the websocket event "BattleStarted" has "payload.id" with value "6b590342-b50d-4f5a-a2d3-f6899014a37b"
