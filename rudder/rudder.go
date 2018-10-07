package rudder

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RudderClient struct {
	endpoint string
	apiToken string
	client   *http.Client

	Nodes NodesClient
}

func NewClient(Endpoint, ApiToken string) *RudderClient {
	client := &RudderClient{
		endpoint: Endpoint,
		apiToken: ApiToken,
	}

	client.Nodes = NodesClient{client: client}

	return client
}

func (client *RudderClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {

	target := client.endpoint + path
	request, err := http.NewRequest(method, target, body)

	if err != nil {
		return nil, err
	}

	//Add Rudder API token to authenticate this request
	request.Header.Add("X-API-Token", client.apiToken)

	return request, nil
}

func (client *RudderClient) Call(req *http.Request) (*Response, error) {

	response, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	//TODO: check for errors: https://docs.rudder.io/api/#api-_-Response - Mostly HTTP 5xx

	//All responses from the API are in the JSON Response format including errors!
	resp := &Response{}
	json.Unmarshal(body, resp)

	return resp, nil
}
