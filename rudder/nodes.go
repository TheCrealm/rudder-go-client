package rudder

type NodesClient struct {
	client *RudderClient
}

type Nodes struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Hostname string `json:"hostname"`
}

//https://docs.rudder.io/api/#api-Nodes-listAcceptedNodes
func (client *NodesClient) ListAcceptedNodes() (*Nodes, error) {

	request, err := client.client.NewRequest("GET", "/api/nodes", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.client.Call(request)
	if err != nil {
		return nil, err
	}

	//get data
	data := &Nodes{}
	response.UnmarschalData(data)

	return data, nil
}
