package entities

type DiscoveredHostGet struct {
	DHostsIDs       []string `json:"dhostsids"`
	DRuleIDs        []string `json:"druleids"`
	DServiceIDs     []string `json:"dserviceids"`
	SelectDRules    string   `json:"selectDRules"`
	SelectDServices string   `json:"selectDServices"`
	LimitSelects    int      `json:"limitSelects"`
	ZabbixCommun
}

type DiscoveredHostOutput struct {
	DServices []DiscoveredServiceOutput `json:"dservices"`
	DRules    []DiscoveredRuleOutput    `json:"drules"`
	DHostID   string                    `json:"dhostid"`
	DRuleID   string                    `json:"druleid"`
	Status    string                    `json:"status"`
	LastUp    string                    `json:"lastup"`
	LastDown  string                    `json:"lastdown"`
}
