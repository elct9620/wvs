package api

import "encoding/json"

var (
	ApiErrInternalServer = json.RawMessage(`{"error":"Internal Server Error"}`)
	ApiErrUnauthorized   = json.RawMessage(`{"error":"Unauthorized"}`)
)
