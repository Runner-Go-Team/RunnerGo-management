package rao

type APIDebug struct {
	ApiID                 int64        `json:"api_id"`
	APIName               string       `json:"api_name"`
	Assertion             []*Assertion `json:"assertion"`
	EventID               string       `json:"event_id"`
	Regex                 string       `json:"regex"`
	RequestBody           string       `json:"request_body"`
	RequestCode           int64        `json:"request_code"`
	RequestHeader         string       `json:"request_header"`
	RequestTime           int64        `json:"request_time"`
	ResponseBody          string       `json:"response_body"`
	ResponseBytes         int64        `json:"response_bytes"`
	ResponseHeader        string       `json:"response_header"`
	ResponseTime          string       `json:"response_time"`
	ResponseLen           int32        `json:"response_len"`
	ResponseStatusMessage string       `json:"response_status_message"`
	UUID                  string       `json:"uuid"`
}

type Assertion struct {
	Code      int    `bson:"code"`
	IsSucceed bool   `bson:"isSucceed"`
	Msg       string `bson:"msg"`
}

type DebugRegex struct {
	Token string `json:"token"`
}
