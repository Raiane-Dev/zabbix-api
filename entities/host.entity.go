package entities

type HostGet struct {
	GroupIDs                      []string `json:"groupids,omitempty"`
	DServiceIDs                   []string `json:"dservicesids,omitempty"`
	GraphIDs                      []string `json:"graphids,omitempty"`
	HostIDs                       []string `json:"hostids,omitempty"`
	HttpTestIDS                   []string `json:"httptestids,omitempty"`
	InterfaceIDs                  []string `json:"interfaceids,omitempty"`
	ItemIDS                       []string `json:"itemids,omitempty"`
	MaintenanceIDS                []string `json:"maintenanceids,omitempty"`
	MonitoredHosts                string   `json:"monitored_hosts,omitempty"`
	ProxyHosts                    string   `json:"proxy_hosts,omitempty"`
	ProxyIDs                      []string `json:"proxyids,omitempty"`
	TemplatedHosts                string   `json:"templated_hosts,omitempty"`
	TemplateIDs                   []string `json:"templateids,omitempty"`
	TriggerIDs                    []string `json:"triggerids,omitempty"`
	WIthItems                     string   `json:"with_items,omitempty"`
	WithItemPrototypes            string   `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes string   `json:"with_simple_graph_item_prototypes,omitempty"`
	WithSimpleGraphPrototypes     string   `json:"with_simple_graph_prototypes,omitempty"`
	WithHttpTests                 string   `json:"with_httptests,omitempty"`
	WithMonitoredHttpTests        string   `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            string   `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers         string   `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems          string   `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  string   `json:"with_triggers,omitempty"`
	WithProblemSupressed          string   `json:"withProblemSupressed,omitempty"`
	Evaltype                      string   `json:"evaltype,omitempty"`
	Severities                    string   `json:"severities,omitempty"`
	Tags                          string   `json:"tags,omitempty"`
	InheritedTags                 string   `json:"inheritedTags,omitempty"`
	SelectDiscoveries             string   `json:"selectDiscoveries,omitempty"`
	SelectDiscoveryRule           string   `json:"selectDiscoveryRule,omitempty"`
	SelectGraphs                  string   `json:"selectGraphs,omitempty"`
	SelectHostDiscovery           string   `json:"selectHostDiscovery,omitempty"`
	SelectHotGroups               string   `json:"selectHotGroups,omitempty"`
	SelectHttpTests               string   `json:"selectHttpTests,omitempty"`
	SelectInterfaces              string   `json:"selectiInterfaces,omitempty"`
	SelectInventory               string   `json:"selectInventory,omitempty"`
	SelectItems                   string   `json:"selectItems,omitempty"`
	SelectMacros                  string   `json:"selectMacros,omitempty"`
	SelectParentTemplates         string   `json:"selectParentTemplates,omitempty"`
	SelectDashboards              string   `json:"selectDashboards,omitempty"`
	SelectTags                    string   `json:"selectTags,omitempty"`
	SelectTriggers                string   `json:"selectTriggers,omitempty"`
	SelectValueMaps               string   `json:"selectValueMaps,omitempty"`
	SelectInheritedTags           string   `json:"selectInheritedTags,omitempty"`
	SearchInventory               string   `json:"searchInventory,omitempty"`
	ZabbixCommon                  string
}

type HostObject struct {
	HostID            string `json:"hostid,omitempty"`
	Host              string `json:"host,omitempty"`
	Description       string `json:"description,omitempty"`
	Flags             int    `json:"flags,omitempty"`
	InventoryMode     int    `json:"inventory_mode,omitempty"`
	IPMIAuthType      int    `json:"ipmi_authtype,omitempty"`
	IPMIPassword      string `json:"ipmi_password,omitempty"`
	IPMIPrivilege     int    `json:"ipmi_privilege,omitempty"`
	IPMIUsername      string `json:"ipmi_username,omitempty"`
	MaintenanceFrom   string `json:"maintenance_from,omitempty"`
	MaintenanceStatus int    `json:"maintenance_status,omitempty"`
	MaintenanceType   int    `json:"maintenance_type,omitempty"`
	MaintenanceID     string `json:"maintenanceid,omitempty"`
	Name              string `json:"name,omitempty"`
	ProxyHostID       string `json:"proxy_hostid,omitempty"`
	Status            int    `json:"status,omitempty"`
	TLSConnect        int    `json:"tls_connect,omitempty"`
	TLSAccept         int    `json:"tls_accept,omitempty"`
	TLSIssuer         string `json:"tls_issuer,omitempty"`
	TLSSubject        string `json:"tls_subject,omitempty"`
	TLSPskIdentify    string `json:"tls_pskidentify,omitempty"`
	TLSPsk            string `json:"tls_psk,omitempty"`
	ActiveAvailable   int    `json:"active_available,omitempty"`
}

type HostCreate struct {
	Host              string                `json:"host,omitempty"`
	Description       string                `json:"description,omitempty"`
	Flags             int                   `json:"flags,omitempty"`
	InventoryMode     int                   `json:"inventory_mode,omitempty"`
	IPMIAuthType      int                   `json:"ipmi_authtype,omitempty"`
	IPMIPassword      string                `json:"ipmi_password,omitempty"`
	IPMIPrivilege     int                   `json:"ipmi_privilege,omitempty"`
	IPMIUsername      string                `json:"ipmi_username,omitempty"`
	MaintenanceFrom   string                `json:"maintenance_from,omitempty"`
	MaintenanceStatus int                   `json:"maintenance_status,omitempty"`
	MaintenanceType   int                   `json:"maintenance_type,omitempty"`
	MaintenanceID     string                `json:"maintenanceid,omitempty"`
	Name              string                `json:"name,omitempty"`
	ProxyHostID       string                `json:"proxy_hostid,omitempty"`
	Status            int                   `json:"status,omitempty"`
	TLSConnect        int                   `json:"tls_connect,omitempty"`
	TLSAccept         int                   `json:"tls_accept,omitempty"`
	TLSIssuer         string                `json:"tls_issuer,omitempty"`
	TLSSubject        string                `json:"tls_subject,omitempty"`
	TLSPskIdentify    string                `json:"tls_pskidentify,omitempty"`
	TLSPsk            string                `json:"tls_psk,omitempty"`
	ActiveAvailable   int                   `json:"active_available,omitempty"`
	Groups            []HostGroupObject     `json:"groups,omitempty"`
	Interfaces        []HostInterfaceObject `json:"interfaces,omitempty"`
	HostTags          []HostTagObject       `json:"hosttags,omitempty"`
	Templates         []TemplateObject      `json:"templates,omitempty"`
	Macros            []UserMacroObject     `json:"macros,omitempty"`
	Inventory         []HostInventoryObject `json:"inventory,omitempty"`
}
type HostUpdate struct {
	HostID            string                `json:"hostid,omitempty"`
	Host              string                `json:"host,omitempty"`
	Description       string                `json:"description,omitempty"`
	Flags             int                   `json:"flags,omitempty"`
	InventoryMode     int                   `json:"inventory_mode,omitempty"`
	IPMIAuthType      int                   `json:"ipmi_authtype,omitempty"`
	IPMIPassword      string                `json:"ipmi_password,omitempty"`
	IPMIPrivilege     int                   `json:"ipmi_privilege,omitempty"`
	IPMIUsername      string                `json:"ipmi_username,omitempty"`
	MaintenanceFrom   string                `json:"maintenance_from,omitempty"`
	MaintenanceStatus int                   `json:"maintenance_status,omitempty"`
	MaintenanceType   int                   `json:"maintenance_type,omitempty"`
	MaintenanceID     string                `json:"maintenanceid,omitempty"`
	Name              string                `json:"name,omitempty"`
	ProxyHostID       string                `json:"proxy_hostid,omitempty"`
	Status            int                   `json:"status,omitempty"`
	TLSConnect        int                   `json:"tls_connect,omitempty"`
	TLSAccept         int                   `json:"tls_accept,omitempty"`
	TLSIssuer         string                `json:"tls_issuer,omitempty"`
	TLSSubject        string                `json:"tls_subject,omitempty"`
	TLSPskIdentify    string                `json:"tls_pskidentify,omitempty"`
	TLSPsk            string                `json:"tls_psk,omitempty"`
	ActiveAvailable   int                   `json:"active_available,omitempty"`
	Groups            []HostGroupObject     `json:"groups,omitempty"`
	Interfaces        []HostInterfaceObject `json:"interfaces,omitempty"`
	HostTags          []HostTagObject       `json:"hosttags,omitempty"`
	Templates         []TemplateObject      `json:"templates,omitempty"`
	Macros            []UserMacroObject     `json:"macros,omitempty"`
	Inventory         []HostInventoryObject `json:"inventory,omitempty"`
}

type HostMassAdd struct {
	Hosts      []map[string]string   `json:"hosts,omitempty"`
	Groups     []HostGroupObject     `json:"groups,omitempty"`
	Interfaces []HostInterfaceObject `json:"interfaces,omitempty"`
	HostTags   []HostTagObject       `json:"hosttags,omitempty"`
	Templates  []TemplateObject      `json:"templates,omitempty"`
	Macros     []UserMacroObject     `json:"macros,omitempty"`
}

type HostMassRemove struct {
	HostIDs          []string              `json:"hostids,omitempty"`
	GroupsIDs        []string              `json:"groupsids,omitempty"`
	Interfaces       []HostInterfaceObject `json:"interfaces,omitempty"`
	Macros           []UserMacroObject     `json:"macros,omitempty"`
	TemplateIDs      []string              `json:"templateids,omitempty"`
	TemplateIDsClear []string              `json:"templateids_clear,omitempty"`
}

type HostMassUpdate struct {
	Hosts      []map[string]string   `json:"hosts,omitempty"`
	Groups     []HostGroupObject     `json:"groups,omitempty"`
	Inventory  []HostInventoryObject `json:"inventory,omitempty"`
	Interfaces []HostInterfaceObject `json:"interfaces,omitempty"`
	Templates  []TemplateObject      `json:"templates,omitempty"`
	Macros     []UserMacroObject     `json:"macros,omitempty"`
}

type HostInventoryObject struct {
	Alias string
}

type HostTagObject struct {
	Tag       string `json:"tag,omitempty"`
	Value     string `json:"value,omitempty"`
	Automatic int    `json:"automatic,omitempty"`
}
