package entities

type HostGroupObject struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
	Flags   int    `json:"flags"`
	UUID    string `json:"uuid"`
}

type HostGroupCreate struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
	Flags   int    `json:"flags"`
}

type HostGroupUpdate struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
	Flags   int    `json:"flags"`
	UUID    string `json:"uuid"`
}

type HostGroupResponse struct {
	GroupIDs []string `json:"groupids"`
}

type HostGroupGet struct {
	GraphIDs                      []string `json:"graphids"`
	GroupIDs                      []string `json:"groupids"`
	HostIDs                       []string `json:"hostids"`
	MaintenanceIDs                []string `json:"maintenanceids"`
	TriggerIDs                    []string `json:"triggerids"`
	WithGraphs                    string   `json:"with_graphs"`
	WithGraphPrototypes           string   `json:"with_graph_prototypes"`
	WithHosts                     string   `json:"with_hosts"`
	WithHttpTests                 string   `json:"with_httptests"`
	WithItems                     string   `json:"with_items"`
	WithItemPrototypes            string   `json:"with_item_prototypes"`
	WithSImpleGraphItemPrototypes string   `json:"with_simple_graph_item_prototypes"`
	WithMonitoredHosts            string   `json:"with_monitored_hosts"`
	WithSimpleGraphItems          string   `json:"with_simple_graph_items"`
	WithTriggers                  string   `json:"with_triggers"`
	SelectDiscoveryRule           string   `json:"selectDiscoveryRule"`
	SelectGroupDiscovery          string   `json:"selectGroupDiscovery"`
	SelectHosts                   string   `json:"selectHosts"`
	ZabbixCommun
}

type HostGroupOutput struct {
	GroupID  string `json:"groupid"`
	Name     string `json:"name"`
	Internal string `json:"internal"`
}

type HostGroupMassAdd struct {
	Groups []map[string]string `json:"groups"`
	Hosts  []map[string]string `json:"hosts"`
}

type HostGroupMassRemove struct {
	GroupIDs []string `json:"groupids"`
	HostIDs  []string `json:"hostids"`
}

type HostGroupMassUpdate struct {
	Groups []map[string]string `json:"groups"`
	Hosts  []map[string]string `json:"hosts"`
}

type HostGroupMassPropagate struct {
	Groups      []map[string]string `json:"groups"`
	Permissions bool                `json:"permissions"`
	TagFilters  bool                `json:"tag_filters"`
}
