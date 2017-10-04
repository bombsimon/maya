package client

const (
	NETWORK_TYPE = "network"
)

type Network struct {
	Resource

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	DefaultPolicyAction string `json:"defaultPolicyAction,omitempty" yaml:"default_policy_action,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Dns []string `json:"dns,omitempty" yaml:"dns,omitempty"`

	DnsSearch []string `json:"dnsSearch,omitempty" yaml:"dns_search,omitempty"`

	HostPorts bool `json:"hostPorts,omitempty" yaml:"host_ports,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	NetworkDriverId string `json:"networkDriverId,omitempty" yaml:"network_driver_id,omitempty"`

	Policy []NetworkPolicyRule `json:"policy,omitempty" yaml:"policy,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Subnets []Subnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type NetworkCollection struct {
	Collection
	Data   []Network `json:"data,omitempty"`
	client *NetworkClient
}

type NetworkClient struct {
	rancherClient *RancherClient
}

type NetworkOperations interface {
	List(opts *ListOpts) (*NetworkCollection, error)
	Create(opts *Network) (*Network, error)
	Update(existing *Network, updates interface{}) (*Network, error)
	ById(id string) (*Network, error)
	Delete(container *Network) error

	ActionCreate(*Network) (*Network, error)

	ActionRemove(*Network) (*Network, error)
}

func newNetworkClient(rancherClient *RancherClient) *NetworkClient {
	return &NetworkClient{
		rancherClient: rancherClient,
	}
}

func (c *NetworkClient) Create(container *Network) (*Network, error) {
	resp := &Network{}
	err := c.rancherClient.doCreate(NETWORK_TYPE, container, resp)
	return resp, err
}

func (c *NetworkClient) Update(existing *Network, updates interface{}) (*Network, error) {
	resp := &Network{}
	err := c.rancherClient.doUpdate(NETWORK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkClient) List(opts *ListOpts) (*NetworkCollection, error) {
	resp := &NetworkCollection{}
	err := c.rancherClient.doList(NETWORK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NetworkCollection) Next() (*NetworkCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NetworkCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NetworkClient) ById(id string) (*Network, error) {
	resp := &Network{}
	err := c.rancherClient.doById(NETWORK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *NetworkClient) Delete(container *Network) error {
	return c.rancherClient.doResourceDelete(NETWORK_TYPE, &container.Resource)
}

func (c *NetworkClient) ActionCreate(resource *Network) (*Network, error) {

	resp := &Network{}

	err := c.rancherClient.doAction(NETWORK_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *NetworkClient) ActionRemove(resource *Network) (*Network, error) {

	resp := &Network{}

	err := c.rancherClient.doAction(NETWORK_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}