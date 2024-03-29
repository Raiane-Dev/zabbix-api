package entities

type HostPrototypeObject struct {
	HostID           string `json:"hostid,omitempty"`
	Host             string `json:"host,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           int    `json:"status,omitempty"`
	InventoryMode    int    `json:"inventory_mode,omitempty"`
	TemplateID       string `json:"templateid,omitempty"`
	Discover         int    `json:"discover,omitempty"`
	CustomInterfaces int    `json:"custom_interfaces,omitempty"`
	UUID             string `json:"uuid,omitempty"`
}

type HostPrototypeCreate struct {
	Host             string                `json:"host,omitempty"`
	Name             string                `json:"name,omitempty"`
	Status           int                   `json:"status,omitempty"`
	InventoryMode    int                   `json:"inventory_mode,omitempty"`
	TemplateID       string                `json:"templateid,omitempty"`
	Discover         int                   `json:"discover,omitempty"`
	CustomInterfaces int                   `json:"custom_interfaces,omitempty"`
	UUID             string                `json:"uuid,omitempty"`
	GroupLinks       []map[string]string   `json:"groupLinks,omitempty"`
	RuleID           string                `json:"ruleid,omitempty"`
	GroupPrototypes  []map[string]string   `json:"groupPrototypes,omitempty"`
	Tags             []ZabbixTag           `json:"tags,omitempty"`
	Interfaces       []HostInterfaceObject `json:"interfaces,omitempty"`
	Templates        []TemplateObject      `json:"templates,omitempty"`
}

type HostPrototypeUpdate struct {
	HostID           string                `json:"hostid,omitempty"`
	Host             string                `json:"host,omitempty"`
	Name             string                `json:"name,omitempty"`
	Status           int                   `json:"status,omitempty"`
	InventoryMode    int                   `json:"inventory_mode,omitempty"`
	TemplateID       string                `json:"templateid,omitempty"`
	Discover         int                   `json:"discover,omitempty"`
	CustomInterfaces int                   `json:"custom_interfaces,omitempty"`
	UUID             string                `json:"uuid,omitempty"`
	GroupLinks       []map[string]string   `json:"groupLinks,omitempty"`
	RuleID           string                `json:"ruleid,omitempty"`
	GroupPrototypes  []map[string]string   `json:"groupPrototypes,omitempty"`
	Tags             []ZabbixTag           `json:"tags,omitempty"`
	Interfaces       []HostInterfaceObject `json:"interfaces,omitempty"`
	Templates        []TemplateObject      `json:"templates,omitempty"`
}

type HostPrototypeResponse struct {
	HostIDs []string `json:"hostids,omitempty"`
}

type HostPrototypeGet struct {
	HostIDs               []string `json:"hostids,omitempty"`
	DiscoveryIDs          []string `json:"discoveryids,omitempty"`
	Inherited             bool     `json:"inherited,omitempty"`
	SelectDiscoveryRule   string   `json:"selectDiscoveryRule,omitempty"`
	SelectInterfaces      string   `json:"selectInterfaces,omitempty"`
	SelectGroupLinks      string   `json:"selectGroupLinks,omitempty"`
	SelectGroupPrototypes string   `json:"selectGroupPrototypes,omitempty"`
	SelectMacros          string   `json:"selectMacros,omitempty"`
	SelectParentHost      string   `json:"selectParentHost,omitempty"`
	SelectTags            string   `json:"selectTags,omitempty"`
	SelectTemplates       string   `json:"selectTemplates,omitempty"`
	ZabbixCommun
}
