package rudder

import (
	"encoding/json"
	"fmt"
)

type NodesClient struct {
	client *RudderClient
}

type Nodes struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Id                          string          `json:"id"`
	Hostname                    string          `json:"hostname"`
	Status                      NodeStatus      `json:"status"`
	ArchitectureDescription     string          `json:"architectureDescription"`
	Description                 string          `json:"description"`
	IpAddresses                 []string        `json:"ipAddresses"`
	LastInventoryDate           string          `json:"lastInventoryDate"`
	Machine                     Machine         `json:"machine"`
	Os                          json.RawMessage `json:"os"`
	ManagementTechnology        json.RawMessage `json:"managementTechnology"`
	PolicyServerId              string          `json:"policyServerId"`
	Properties                  json.RawMessage `json:"properties"`
	PolicyMode                  string          `json:"policyMode"`
	Ram                         int             `json:"ram"`
	Timezone                    json.RawMessage `json:"timezone"`
	Accounts                    json.RawMessage `json:"accounts"`
	Bios                        json.RawMessage `json:"bios"`
	Controllers                 json.RawMessage `json:"controllers"`
	EnvironmentVariables        json.RawMessage `json:"environmentVariables"`
	FileSystems                 json.RawMessage `json:"fileSystems"`
	ManagementTechnologyDetails json.RawMessage `json:"managementTechnologyDetails"`
	Memories                    json.RawMessage `json:"memories"`
	NetworkInterfaces           json.RawMessage `json:"networkInterfaces"`
	Ports                       json.RawMessage `json:"ports"`
	Processes                   json.RawMessage `json:"processes"`
	Processors                  json.RawMessage `json:"processors"`
	Slots                       json.RawMessage `json:"slots"`
	Software                    json.RawMessage `json:"software"`
	Sound                       json.RawMessage `json:"sound"`
	Storage                     json.RawMessage `json:"storage"`
	Videos                      json.RawMessage `json:"videos"`
	VirtualMachines             json.RawMessage `json:"virtualMachines"`
}

type NodeStatus string

var (
	Pending  NodeStatus = "pending"
	Accepted NodeStatus = "accepted"
	Deleted  NodeStatus = "deleted"
)

type Machine struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Provider     string `json:"provider"`
	Manufacturer string `json:"manufacturer"`
	SerialNumber string `json:"serialNumber"`
}

type QueryWhere struct {
	ObjectType string `json:"objectType"`
	Attribute  string `json:"attribute"`
	Comperator string `json:"comperator"`
	Value      string `json:"value"`
}

//https://docs.rudder.io/api/#api-Nodes-listAcceptedNodes
func (client *NodesClient) ListAcceptedNodes() (*Nodes, error) {

	data := &Nodes{}

	qwhere := &QueryWhere{
		ObjectType: "node",
		Attribute:  "OS",
		Comperator: "eq",
		Value:      "Linux",
	}

	_, err := json.Marshal(qwhere)
	if err != nil {
		return nil, err
	}

	//TODO: Generate Node Query - this query is just for testing purposes here
	query := `?include=full&where=[{"objectType":"node","attribute":"OS","comparator":"eq","value":"Linux"}]`

	_, err = client.client.Execute("GET", fmt.Sprint("/api/nodes", query), nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (client *NodesClient) ListPendingNodes() (*Nodes, error) {

	data := &Nodes{}
	_, err := client.client.Execute("GET", "/api/nodes/pending", nil, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
