package entities

type HighAvailabilityNodeGet struct {
	HaNodeIDs []string `json:"ha_nodeids,omitempty"`
	ZabbixCommun
}

type HighAvailabilityNodeOutput struct {
	HaNodeID   string `json:"ha_nodeid"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	LastAccess string `json:"lastaccess"`
	Status     string `json:"status"`
}
