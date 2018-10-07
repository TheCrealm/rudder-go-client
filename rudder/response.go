package rudder

import "encoding/json"

//https://docs.rudder.io/api/#api-_-Response
type Response struct {
	Action       string          `json:"action"`
	Id           string          `json:"id"`
	Result       string          `json:"result"`
	Data         json.RawMessage `json:"data"`
	ErrorDetails string          `json:"errorDetails"`
}

func (resp *Response) UnmarschalData(v interface{}) {
	json.Unmarshal([]byte(resp.Data), v)
}
