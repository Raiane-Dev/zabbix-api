package entities

type DiscoveredRuleCreate struct {
	DChecks     []DiscoveredCheckOutput
	IPRange     string `json:"iprange"`
	Name        string `json:"name"`
	Delay       string `json:"delay"`
	ProxyHostID string `json:"proxy_hostid"`
	Status      int    `json:"status"`
}

type DiscoveredRuleUpdate struct {
	DRuleID     string `json:"druleid"`
	DChecks     []DiscoveredCheckOutput
	IPRange     string `json:"iprange"`
	Name        string `json:"name"`
	Delay       string `json:"delay"`
	ProxyHostID string `json:"proxy_hostid"`
	Status      int    `json:"status"`
}

type DiscoveredRuleResponse struct {
	DRuleIDs []string `json:"druleids"`
}

type DiscoveredRuleGet struct {
	DHostIDs      []string `json:"dhostids"`
	DRuleIDs      []string `json:"druleids"`
	DServiceIDs   []string `json:"dserviceids"`
	SelectDChecks string   `json:"selectDChecks"`
	SelectDHosts  string   `json:"selectDHosts"`
	LimitSelects  int      `json:"limitSelects"`
	ZabbixCommun
}

type DiscoveredRuleOutput struct {
	DRuleID     string `json:"druleid"`
	IPRange     string `json:"iprange"`
	Name        string `json:"name"`
	Delay       string `json:"delay"`
	ProxyHostID string `json:"proxy_hostid"`
	Status      int    `json:"status"`
}
