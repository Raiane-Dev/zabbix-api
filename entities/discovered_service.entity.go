package entities

type DiscoveredServiceGet struct {
	DServiceIDs  []string `json:"dserviceid,omitempty"`
	DHostIDs     []string `json:"dhostid,omitempty"`
	DCheckIDs    []string `json:"dcheckids,omitempty"`
	DRuleIDs     []string `json:"druleids,omitempty"`
	SelectDRules string   `json:"selectDRules,omitempty"`
	SelectDHosts string   `json:"selectDHosts,omitempty"`
	SelectHosts  string   `json:"selectHosts,omitempty"`
	LimitSelects int      `json:"limitSelects,omitempty"`
	ZabbixCommun
}

type DiscoveredServiceObject struct {
	DServiceIDs string `json:"dserviceids,omitempty"`
	DHostID     string `json:"dhostid,omitempty"`
	Type        string `json:"type,omitempty"`
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	Port        string `json:"port,omitempty"`
	Status      int    `json:"status,omitempty"`
	LastUp      string `json:"lastup,omitempty"`
	LastDown    string `json:"lastdown,omitempty"`
	DCheckID    string `json:"dcheckid,omitempty"`
	IP          string `json:"ip,omitempty"`
	DNS         string `json:"dns,omitempty"`
}
