package rudder

import "fmt"

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

	query := `?where=[{"objectType":"node","attribute":"OS","comparator":"eq","value":"Linux"}]`
	request, err := client.client.NewRequest("GET", fmt.Sprint("/api/nodes", query), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(request.URL)

	response, err := client.client.Call(request)
	if err != nil {
		return nil, err
	}

	//get data
	data := &Nodes{}
	response.UnmarschalData(data)

	return data, nil
}
