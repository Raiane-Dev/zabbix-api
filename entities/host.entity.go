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
	ZabbixCommon                  string   `json:"-"`
}

type HostObject struct {
	HostID            string `json:"hostid,omitempty"`
	Host              string `json:"host"`
	Description       string `json:"description"`
	Flags             int    `json:"flags"`
	InventoryMode     int    `json:"inventory_mode"`
	IPMIAuthType      int    `json:"ipmi_authtype"`
	IPMIPassword      string `json:"ipmi_password"`
	IPMIPrivilege     int    `json:"ipmi_privilege"`
	IPMIUsername      string `json:"ipmi_username"`
	MaintenanceFrom   string `json:"maintenance_from"`
	MaintenanceStatus int    `json:"maintenance_status"`
	MaintenanceType   int    `json:"maintenance_type"`
	MaintenanceID     string `json:"maintenanceid"`
	Name              string `json:"name"`
	ProxyHostID       string `json:"proxy_hostid"`
	Status            int    `json:"status"`
	TLSConnect        int    `json:"tls_connect"`
	TLSAccept         int    `json:"tls_accept"`
	TLSIssuer         string `json:"tls_issuer"`
	TLSSubject        string `json:"tls_subject"`
	TLSPskIdentify    string `json:"tls_pskidentify"`
	TLSPsk            string `json:"tls_psk"`
	ActiveAvailable   int    `json:"active_available"`
}

type HostCreate struct {
	Host              string                `json:"host"`
	Description       string                `json:"description"`
	Flags             int                   `json:"flags"`
	InventoryMode     int                   `json:"inventory_mode"`
	IPMIAuthType      int                   `json:"ipmi_authtype"`
	IPMIPassword      string                `json:"ipmi_password"`
	IPMIPrivilege     int                   `json:"ipmi_privilege"`
	IPMIUsername      string                `json:"ipmi_username"`
	MaintenanceFrom   string                `json:"maintenance_from"`
	MaintenanceStatus int                   `json:"maintenance_status"`
	MaintenanceType   int                   `json:"maintenance_type"`
	MaintenanceID     string                `json:"maintenanceid"`
	Name              string                `json:"name"`
	ProxyHostID       string                `json:"proxy_hostid"`
	Status            int                   `json:"status"`
	TLSConnect        int                   `json:"tls_connect"`
	TLSAccept         int                   `json:"tls_accept"`
	TLSIssuer         string                `json:"tls_issuer"`
	TLSSubject        string                `json:"tls_subject"`
	TLSPskIdentify    string                `json:"tls_pskidentify"`
	TLSPsk            string                `json:"tls_psk"`
	ActiveAvailable   int                   `json:"active_available"`
	Groups            []HostGroupObject     `json:"groups"`
	Interfaces        []HostInterfaceObject `json:"interfaces"`
	HostTags          []HostTagObject       `json:"hosttags"`
	Templates         []TemplateObject      `json:"templates"`
	Macros            []UserMacroObject     `json:"macros"`
	Inventory         []HostInventoryObject `json:"inventory"`
}
type HostUpdate struct {
	HostID            string                `json:"hostid,omitempty"`
	Host              string                `json:"host"`
	Description       string                `json:"description"`
	Flags             int                   `json:"flags"`
	InventoryMode     int                   `json:"inventory_mode"`
	IPMIAuthType      int                   `json:"ipmi_authtype"`
	IPMIPassword      string                `json:"ipmi_password"`
	IPMIPrivilege     int                   `json:"ipmi_privilege"`
	IPMIUsername      string                `json:"ipmi_username"`
	MaintenanceFrom   string                `json:"maintenance_from"`
	MaintenanceStatus int                   `json:"maintenance_status"`
	MaintenanceType   int                   `json:"maintenance_type"`
	MaintenanceID     string                `json:"maintenanceid"`
	Name              string                `json:"name"`
	ProxyHostID       string                `json:"proxy_hostid"`
	Status            int                   `json:"status"`
	TLSConnect        int                   `json:"tls_connect"`
	TLSAccept         int                   `json:"tls_accept"`
	TLSIssuer         string                `json:"tls_issuer"`
	TLSSubject        string                `json:"tls_subject"`
	TLSPskIdentify    string                `json:"tls_pskidentify"`
	TLSPsk            string                `json:"tls_psk"`
	ActiveAvailable   int                   `json:"active_available"`
	Groups            []HostGroupObject     `json:"groups"`
	Interfaces        []HostInterfaceObject `json:"interfaces"`
	HostTags          []HostTagObject       `json:"hosttags"`
	Templates         []TemplateObject      `json:"templates"`
	Macros            []UserMacroObject     `json:"macros"`
	Inventory         []HostInventoryObject `json:"inventory"`
}

type HostMassAdd struct {
	Hosts      []map[string]string   `json:"hosts"`
	Groups     []HostGroupObject     `json:"groups"`
	Interfaces []HostInterfaceObject `json:"interfaces"`
	HostTags   []HostTagObject       `json:"hosttags"`
	Templates  []TemplateObject      `json:"templates"`
	Macros     []UserMacroObject     `json:"macros"`
}

type HostMassRemove struct {
	HostIDs          []string              `json:"hostids"`
	GroupsIDs        []string              `json:"groupsids"`
	Interfaces       []HostInterfaceObject `json:"interfaces"`
	Macros           []UserMacroObject     `json:"macros"`
	TemplateIDs      []string              `json:"templateids"`
	TemplateIDsClear []string              `json:"templateids_clear"`
}

type HostMassUpdate struct {
	Hosts      []map[string]string   `json:"hosts"`
	Groups     []HostGroupObject     `json:"groups"`
	Inventory  []HostInventoryObject `json:"inventory"`
	Interfaces []HostInterfaceObject `json:"interfaces"`
	Templates  []TemplateObject      `json:"templates"`
	Macros     []UserMacroObject     `json:"macros"`
}

type HostInventoryObject struct {
	Alias string
}

type HostTagObject struct {
	Tag       string `json:"tag"`
	Value     string `json:"value"`
	Automatic int    `json:"automatic"`
}
