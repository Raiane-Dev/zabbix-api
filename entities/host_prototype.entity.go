package entities

type HostPrototypeObject struct {
	HostID           string `json:"hostid"`
	Host             string `json:"host"`
	Name             string `json:"name"`
	Status           int    `json:"status"`
	InventoryMode    int    `json:"inventory_mode"`
	TemplateID       string `json:"templateid"`
	Discover         int    `json:"discover"`
	CustomInterfaces int    `json:"custom_interfaces"`
	UUID             string `json:"uuid"`
}

type HostPrototypeCreate struct {
	Host             string                `json:"host"`
	Name             string                `json:"name"`
	Status           int                   `json:"status"`
	InventoryMode    int                   `json:"inventory_mode"`
	TemplateID       string                `json:"templateid"`
	Discover         int                   `json:"discover"`
	CustomInterfaces int                   `json:"custom_interfaces"`
	UUID             string                `json:"uuid"`
	GroupLinks       []map[string]string   `json:"groupLinks"`
	RuleID           string                `json:"ruleid"`
	GroupPrototypes  []map[string]string   `json:"groupPrototypes"`
	Tags             []ZabbixTag           `json:"tags"`
	Interfaces       []HostInterfaceObject `json:"interfaces"`
	Templates        []TemplateObject      `json:"templates"`
}

type HostPrototypeUpdate struct {
	HostID           string                `json:"hostid"`
	Host             string                `json:"host"`
	Name             string                `json:"name"`
	Status           int                   `json:"status"`
	InventoryMode    int                   `json:"inventory_mode"`
	TemplateID       string                `json:"templateid"`
	Discover         int                   `json:"discover"`
	CustomInterfaces int                   `json:"custom_interfaces"`
	UUID             string                `json:"uuid"`
	GroupLinks       []map[string]string   `json:"groupLinks"`
	RuleID           string                `json:"ruleid"`
	GroupPrototypes  []map[string]string   `json:"groupPrototypes"`
	Tags             []ZabbixTag           `json:"tags"`
	Interfaces       []HostInterfaceObject `json:"interfaces"`
	Templates        []TemplateObject      `json:"templates"`
}

type HostPrototypeResponse struct {
	HostIDs []string `json:"hostids"`
}

type HostPrototypeGet struct {
	HostIDs               []string `json:"hostids"`
	DiscoveryIDs          []string `json:"discoveryids"`
	Inherited             bool     `json:"inherited"`
	SelectDiscoveryRule   string   `json:"selectDiscoveryRule"`
	SelectInterfaces      string   `json:"selectInterfaces"`
	SelectGroupLinks      string   `json:"selectGroupLinks"`
	SelectGroupPrototypes string   `json:"selectGroupPrototypes"`
	SelectMacros          string   `json:"selectMacros"`
	SelectParentHost      string   `json:"selectParentHost"`
	SelectTags            string   `json:"selectTags"`
	SelectTemplates       string   `json:"selectTemplates"`
	ZabbixCommun
}
