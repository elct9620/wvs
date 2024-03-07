Feature: WebSocket
  Scenario: when connected the ready event is received
    When connect to the websocket
    Then the websocket event is received
      """
      {
        "event": "ready"
      }
      """
