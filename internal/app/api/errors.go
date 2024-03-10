package api

import "encoding/json"

var (
	ApiErrInternalServer   = json.RawMessage(`{"error":"Internal Server Error"}`)
	ApiErrUnauthorized     = json.RawMessage(`{"error":"Unauthorized"}`)
	ApiErrUnableReadBody   = json.RawMessage(`{"error":"Unable to read request body"}`)
	ApiErrDecodeJsonFailed = json.RawMessage(`{"error":"Failed to decode JSON"}`)
)
