package entities

type DiscoveredServiceGet struct {
	DServiceIDs  []string `json:"dserviceid"`
	DHostIDs     []string `json:"dhostid"`
	DCheckIDs    []string `json:"dcheckids"`
	DRuleIDs     []string `json:"druleids"`
	SelectDRules string   `json:"selectDRules"`
	SelectDHosts string   `json:"selectDHosts"`
	SelectHosts  string   `json:"selectHosts"`
	LimitSelects int      `json:"limitSelects"`
	ZabbixCommun
}

type DiscoveredServiceOutput struct {
	DServiceIDs string
	DHostID     string
	Type        string
	Key         string
	Value       string
	Port        string
	Status      int
	LastUp      string
	LastDown    string
	DCheckID    string
	IP          string
	DNS         string
}
