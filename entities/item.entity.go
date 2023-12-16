package entities

type ItemGet struct {
	ItemIDs             []string            `json:"itemids,omitempty"`
	GroupIDs            []string            `json:"groupids,omitempty"`
	TemplateIDs         []string            `json:"templateids,omitempty"`
	HostIDs             []string            `json:"hostids,omitempty"`
	ProxyIDs            []string            `json:"proxyids,omitempty"`
	InterfaceIDs        []string            `json:"interfaceids,omitempty"`
	GraphIDS            []string            `json:"graphids,omitempty"`
	TriggerIDs          []string            `json:"triggerids,omitempty"`
	WebItems            string              `json:"webitems,omitempty"`
	Inherited           bool                `json:"inherited,omitempty"`
	Templated           bool                `json:"templated,omitempty"`
	Monitored           bool                `json:"monitored,omitempty"`
	Group               string              `json:"group,omitempty"`
	Host                string              `json:"host,omitempty"`
	EvalType            int                 `json:"evaltype,omitempty"`
	Tags                []map[string]string `json:"tags,omitempty"`
	WithTriggers        bool                `json:"with_triggers,omitempty"`
	SelectHosts         string              `json:"selectHosts,omitempty"`
	SelectInterfaces    string              `json:"selectInterfaces,omitempty"`
	SelectTriggers      string              `json:"selectTriggers,omitempty"`
	SelectGraphs        string              `json:"selectGraphs,omitempty"`
	SelectTags          string              `json:"selectTags,omitempty"`
	SelectDiscoveryRule string              `json:"selectDiscoveryRule,omitempty"`
	SelectItemDiscovery string              `json:"selectItemDiscovery,omitempty"`
	SelectPreprocessing string              `json:"selectPreprocessing,omitempty"`
	ZabbixCommun
}

type ItemCreate struct {
	ItemIDs             []string            `json:"itemids,omitempty"`
	GroupIDs            []string            `json:"groupids,omitempty"`
	TemplateIDs         []string            `json:"templateids,omitempty"`
	HostIDs             []string            `json:"hostids,omitempty"`
	ProxyIDs            []string            `json:"proxyids,omitempty"`
	InterfaceIDs        []string            `json:"interfaceids,omitempty"`
	GraphIDS            []string            `json:"graphids,omitempty"`
	TriggerIDs          []string            `json:"triggerids,omitempty"`
	WebItems            string              `json:"webitems,omitempty"`
	Inherited           bool                `json:"inherited,omitempty"`
	Templated           bool                `json:"templated,omitempty"`
	Monitored           bool                `json:"monitored,omitempty"`
	Group               string              `json:"group,omitempty"`
	Host                string              `json:"host,omitempty"`
	EvalType            int                 `json:"evaltype,omitempty"`
	Tags                []map[string]string `json:"tags,omitempty"`
	WithTriggers        bool                `json:"with_triggers,omitempty"`
	SelectHosts         string              `json:"selectHosts,omitempty"`
	SelectInterfaces    string              `json:"selectInterfaces,omitempty"`
	SelectTriggers      string              `json:"selectTriggers,omitempty"`
	SelectGraphs        string              `json:"selectGraphs,omitempty"`
	SelectTags          string              `json:"selectTags,omitempty"`
	SelectDiscoveryRule string              `json:"selectDiscoveryRule,omitempty"`
	SelectItemDiscovery string              `json:"selectItemDiscovery,omitempty"`
	SelectPreprocessing string              `json:"selectPreprocessing,omitempty"`
}

type ItemUpdate struct {
}

type ItemDelete struct {
}

type ItemOutput struct {
	ItemID      string              `json:"itemid"`
	Type        string              `json:"type"`
	HostID      string              `json:"hostid"`
	Name        string              `json:"name"`
	Key         string              `json:"key_"`
	Delay       string              `json:"delay"`
	Tags        []map[string]string `json:"tags"`
	History     string              `json:"history"`
	Trends      string              `json:"trends"`
	Status      string              `json:"status"`
	Description string              `json:"description"`
	Timeout     string              `json:"timeout"`
	Lastclock   string              `json:"lastclock"`
	Lastns      string              `json:"lastns"`
	Lastvalue   string              `json:"lastvalue"`
	Prevvalue   string              `json:"prevvalue"`
}
