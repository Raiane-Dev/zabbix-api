package entities

type DiscoveryRuleCreate struct {
	DChecks     []DiscoveryCheckObject `json:"dchecks,omitempty"`
	IPRange     string                 `json:"iprange,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Delay       string                 `json:"delay,omitempty"`
	ProxyHostID string                 `json:"proxy_hostid,omitempty"`
	Status      int                    `json:"status,omitempty"`
}

type DiscoveryRuleUpdate struct {
	DChecks     []DiscoveryCheckObject `json:"dchecks,omitempty"`
	DRuleID     string                 `json:"druleid,omitempty"`
	IPRange     string                 `json:"iprange,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Delay       string                 `json:"delay,omitempty"`
	ProxyHostID string                 `json:"proxy_hostid,omitempty"`
	Status      int                    `json:"status,omitempty"`
}

type DiscoveryRuleResponse struct {
	DRuleIDs []string `json:"druleids,omitempty"`
}

type DiscoveryRuleGet struct {
	DHostIDs      []string `json:"dhostids,omitempty"`
	DRuleIDs      []string `json:"druleids,omitempty"`
	DServiceIDs   []string `json:"dserviceids,omitempty"`
	SelectDChecks string   `json:"selectDChecks,omitempty"`
	SelectDHosts  string   `json:"selectDHosts,omitempty"`
	LimitSelects  int      `json:"limitSelects,omitempty"`
	ZabbixCommun
}

type DiscoveryRuleObject struct {
	DRuleID     string `json:"druleid,omitempty"`
	IPRange     string `json:"iprange,omitempty"`
	Name        string `json:"name,omitempty"`
	Delay       string `json:"delay,omitempty"`
	ProxyHostID string `json:"proxy_hostid,omitempty"`
	Status      int    `json:"status,omitempty"`
}
