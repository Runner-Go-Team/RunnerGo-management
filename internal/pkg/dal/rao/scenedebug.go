package rao

type SceneDebug struct {
	ApiID          int64             `json:"api_id"`
	APIName        string            `json:"api_name"`
	Assertion      []*DebugAssertion `json:"assertion"`
	EventID        string            `json:"event_id"`
	NextList       []string          `json:"next_list"`
	Regex          []*DebugRegex     `json:"regex"`
	RequestBody    string            `json:"request_body"`
	RequestCode    int64             `json:"request_code"`
	RequestHeader  string            `json:"request_header"`
	ResponseBody   string            `json:"response_body"`
	ResponseBytes  int64             `json:"response_bytes"`
	ResponseHeader string            `json:"response_header"`
	Status         string            `json:"status"`
	UUID           string            `json:"uuid"`
}