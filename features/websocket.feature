Feature: WebSocket
  Scenario: when connected the ready event is received
    Given the session id is "d7ae7356-2d91-47f7-81bd-428c40bf55c3"
    When connect to the websocket
    Then the websocket event "ReadyEvent" is received
