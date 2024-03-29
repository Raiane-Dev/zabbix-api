package entities

type HighAvailabilityNodeGet struct {
	HaNodeIDs []string `json:"ha_nodeids,omitempty"`
	ZabbixCommun
}

type HighAvailabilityNodeObject struct {
	HaNodeID   string `json:"ha_nodeid,omitempty"`
	Name       string `json:"name,omitempty"`
	Address    string `json:"address,omitempty"`
	Port       string `json:"port,omitempty"`
	LastAccess string `json:"lastaccess,omitempty"`
	Status     string `json:"status,omitempty"`
}
