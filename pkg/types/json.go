package types

// Backend defines any backend with its properties.
type Backend struct {
	Name     string `json:"name"`
	IPAddr   string `json:"ip-addr"`
	Weight   string `json:"weight,omitempty"`
	Priority string `json:"priority,omitempty"`
	Mark     string `json:"mark,omitempty"`
	State    string `json:"state,omitempty"`
}

// Backends defines a group of backends in any farm.
type Backends []Backend

// Farm defines any farm with its properties.
type Farm struct {
	Name         string   `json:"name"`
	Iface        string   `json:"iface,omitempty"`
	Oface        string   `json:"oface,omitempty"`
	Family       string   `json:"family,omitempty"`
	EtherAddr    string   `json:"ether-addr,omitempty"`
	VirtualAddr  string   `json:"virtual-addr,omitempty"`
	VirtualPorts string   `json:"virtual-ports,omitempty"`
	Mode         string   `json:"mode,omitempty"`
	Protocol     string   `json:"protocol,omitempty"`
	Scheduler    string   `json:"scheduler,omitempty"`
	Helper       string   `json:"helper,omitempty"`
	Log          string   `json:"log,omitempty"`
	Mark         string   `json:"mark,omitempty"`
	Priority     string   `json:"priority,omitempty"`
	State        string   `json:"state,omitempty"`
	Backends     Backends `json:"backends"`
}

// Farms defines a group of farms.
type Farms []Farm

// JSONnftlb is a JSON object made for nftlb requests.
type JSONnftlb struct {
	Farms Farms `json:"farms"`
}
