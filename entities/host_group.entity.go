package entities

type HostGroupObject struct {
	GroupID  string `json:"groupid,omitempty"`
	Name     string `json:"name,omitempty"`
	Flags    int    `json:"flags,omitempty"`
	UUID     string `json:"uuid,omitempty"`
	Internal string `json:"internal,omitempty"`
}

type HostGroupCreate struct {
	GroupID string `json:"groupid,omitempty"`
	Name    string `json:"name,omitempty"`
	Flags   int    `json:"flags,omitempty"`
}

type HostGroupUpdate struct {
	GroupID string `json:"groupid,omitempty"`
	Name    string `json:"name,omitempty"`
	Flags   int    `json:"flags,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type HostGroupResponse struct {
	GroupIDs []string `json:"groupids,omitempty"`
}

type HostGroupGet struct {
	GraphIDs                      []string `json:"graphids,omitempty"`
	GroupIDs                      []string `json:"groupids,omitempty"`
	HostIDs                       []string `json:"hostids,omitempty"`
	MaintenanceIDs                []string `json:"maintenanceids,omitempty"`
	TriggerIDs                    []string `json:"triggerids,omitempty"`
	WithGraphs                    string   `json:"with_graphs,omitempty"`
	WithGraphPrototypes           string   `json:"with_graph_prototypes,omitempty"`
	WithHosts                     string   `json:"with_hosts,omitempty"`
	WithHttpTests                 string   `json:"with_httptests,omitempty"`
	WithItems                     string   `json:"with_items,omitempty"`
	WithItemPrototypes            string   `json:"with_item_prototypes,omitempty"`
	WithSImpleGraphItemPrototypes string   `json:"with_simple_graph_item_prototypes,omitempty"`
	WithMonitoredHosts            string   `json:"with_monitored_hosts,omitempty"`
	WithSimpleGraphItems          string   `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  string   `json:"with_triggers,omitempty"`
	SelectDiscoveryRule           string   `json:"selectDiscoveryRule,omitempty"`
	SelectGroupDiscovery          string   `json:"selectGroupDiscovery,omitempty"`
	SelectHosts                   string   `json:"selectHosts,omitempty"`
	ZabbixCommun
}

type HostGroupMassAdd struct {
	Groups []map[string]string `json:"groups,omitempty"`
	Hosts  []map[string]string `json:"hosts,omitempty"`
}

type HostGroupMassRemove struct {
	GroupIDs []string `json:"groupids,omitempty"`
	HostIDs  []string `json:"hostids,omitempty"`
}

type HostGroupMassUpdate struct {
	Groups []map[string]string `json:"groups,omitempty"`
	Hosts  []map[string]string `json:"hosts,omitempty"`
}

type HostGroupMassPropagate struct {
	Groups      []map[string]string `json:"groups,omitempty"`
	Permissions bool                `json:"permissions,omitempty"`
	TagFilters  bool                `json:"tag_filters,omitempty"`
}
