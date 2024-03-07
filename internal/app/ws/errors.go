package ws

import "encoding/json"

var (
	WsErrUpgrading = json.RawMessage(`{"error":"could not upgrade"}`)
)
