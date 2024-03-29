package entities

type DiscoveredHostGet struct {
	DHostsIDs       []string `json:"dhostsids,,omitempty"`
	DRuleIDs        []string `json:"druleids,,omitempty"`
	DServiceIDs     []string `json:"dserviceids,,omitempty"`
	SelectDRules    string   `json:"selectDRules,,omitempty"`
	SelectDServices string   `json:"selectDServices,,omitempty"`
	LimitSelects    int      `json:"limitSelects,,omitempty"`
	ZabbixCommun
}

type DiscoveredHostObject struct {
	DServices []DiscoveredServiceObject `json:"dservices,omitempty"`
	DRules    []DiscoveryRuleObject     `json:"drules,omitempty"`
	DHostID   string                    `json:"dhostid,omitempty"`
	DRuleID   string                    `json:"druleid,omitempty"`
	Status    string                    `json:"status,omitempty"`
	LastUp    string                    `json:"lastup,omitempty"`
	LastDown  string                    `json:"lastdown,omitempty"`
}
