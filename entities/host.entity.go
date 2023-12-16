package entities

type HostGet struct {
	GroupIDs                      []string          `json:"groupids,omitempty"`
	DServiceIDs                   []string          `json:"dservicesids,omitempty"`
	GraphIDs                      []string          `json:"graphids,omitempty"`
	HostIDs                       []string          `json:"hostids,omitempty"`
	HttpTestIDS                   []string          `json:"httptestids,omitempty"`
	InterfaceIDs                  []string          `json:"interfaceids,omitempty"`
	ItemIDS                       []string          `json:"itemids,omitempty"`
	MaintenanceIDS                []string          `json:"maintenanceids,omitempty"`
	MonitoredHosts                string            `json:"monitored_hosts,omitempty"`
	ProxyHosts                    string            `json:"proxy_hosts,omitempty"`
	ProxyIDs                      []string          `json:"proxyids,omitempty"`
	TemplatedHosts                string            `json:"templated_hosts,omitempty"`
	TemplateIDs                   []string          `json:"templateids,omitempty"`
	TriggerIDs                    []string          `json:"triggerids,omitempty"`
	WIthItems                     string            `json:"with_items,omitempty"`
	WithItemPrototypes            string            `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes string            `json:"with_simple_graph_item_prototypes,omitempty"`
	WithSimpleGraphPrototypes     string            `json:"with_simple_graph_prototypes,omitempty"`
	WithHttpTests                 string            `json:"with_httptests,omitempty"`
	WithMonitoredHttpTests        string            `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            string            `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers         string            `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems          string            `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  string            `json:"with_triggers,omitempty"`
	WithProblemSupressed          string            `json:"withProblemSupressed,omitempty"`
	Evaltype                      string            `json:"evaltype,omitempty"`
	Severities                    string            `json:"severities,omitempty"`
	Tags                          string            `json:"tags,omitempty"`
	InheritedTags                 string            `json:"inheritedTags,omitempty"`
	SelectDiscoveries             string            `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule           string            `json:"selectDiscoveryRule,omitempty"`
	SelectGraphs                  string            `json:"selectGraphs,omitempty"`
	SelectHostDiscovery           string            `json:"selectHostDiscovery,omitempty"`
	SelectHotGroups               string            `json:"selectHotGroups,omitempty"`
	SelectHttpTests               string            `json:"selectHttpTests,omitempty"`
	SelectInterfaces              string            `json:"selectiInterfaces,omitempty"`
	SelectInventory               string            `json:"selectInventory,omitempty"`
	SelectItems                   string            `json:"selectItems,omitempty"`
	SelectMacros                  string            `json:"selectMacros,omitempty"`
	SelectParentTemplates         string            `json:"selectParentTemplates,omitempty"`
	SelectDashboards              string            `json:"selectDashboards,omitempty"`
	SelectTags                    string            `json:"selectTags,omitempty"`
	SelectTriggers                string            `json:"selectTriggers,omitempty"`
	SelectValueMaps               string            `json:"selectValueMaps,omitempty"`
	Filter                        string            `json:"filter,omitempty"`
	SelectInheritedTags           string            `json:"selectInheritedTags,omitempty"`
	LimitSelects                  string            `json:"limitSelects,omitempty"`
	Search                        string            `json:"search,omitempty"`
	SearchInventory               string            `json:"searchInventory,omitempty"`
	Sortfield                     map[string]string `json:"sortfield,omitempty"`
	ZabbixCommon                  string            `json:"-"`
}

type HostOutput struct {
	HostID            string `json:"hostid,omitempty"`
	ProxyHostID       string `json:"proxy_hostid,omitempty"`
	Host              string `json:"host,omitempty"`
	Status            string `json:"status,omitempty"`
	LastAccess        string `json:"lastaccess,omitempty"`
	IPMIAuthType      string `json:"ipmi_authtype,omitempty"`
	IPMIPrivilege     string `json:"ipmi_privilege,omitempty"`
	IPMIUsername      string `json:"ipmi_username,omitempty"`
	IPMIPassword      string `json:"ipmi_password,omitempty"`
	MaintenanceID     string `json:"maintenanceid,omitempty"`
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
	MaintenanceType   string `json:"maintenance_type,omitempty"`
	MaintenanceFrom   string `json:"maintenance_from,omitempty"`
	Name              string `json:"name,omitempty"`
	Flags             string `json:"flags,omitempty"`
	Description       string `json:"description,omitempty"`
	TLSConnect        string `json:"tls_connect,omitempty"`
	TLSAccept         string `json:"tls_accept,omitempty"`
	TLSIssuer         string `json:"tls_issuer,omitempty"`
	TLSSubject        string `json:"tls_subject,omitempty"`
	InventoryMode     string `json:"inventory_mode,omitempty"`
	ActiveAvailable   string `json:"active_available,omitempty"`
}

type HostCreate struct {
}

type HostUpdate struct {
}

type HostDelete struct {
}
