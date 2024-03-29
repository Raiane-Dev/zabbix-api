package entities

type TemplateObject struct {
	TemplateID     string `json:"templateid,omitempty"`
	Host           string `json:"host,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	UUID           string `json:"uuid,omitempty"`
	VendorName     string `json:"vendor_name,omitempty"`
	VendorVersionn string `json:"vendor_versionn,omitempty"`
}

type TemplateTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

type TemplateCreate struct {
	Host           string                `json:"host,omitempty"`
	Description    string                `json:"description,omitempty"`
	Name           string                `json:"name,omitempty"`
	UUID           string                `json:"uuid,omitempty"`
	VendorName     string                `json:"vendor_name,omitempty"`
	VendorVersionn string                `json:"vendor_versionn,omitempty"`
	Groups         []TemplateGroupObject `json:"groups,omitempty"`
	Tags           []TemplateTag         `json:"tags,omitempty"`
	Templates      []TemplateObject      `json:"templates,omitempty"`
	Macros         []UserMacroObject     `json:"macros,omitempty"`
}

type TemplateUpdate struct {
	TemplateID     string                `json:"templateid,omitempty"`
	Host           string                `json:"host,omitempty"`
	Description    string                `json:"description,omitempty"`
	Name           string                `json:"name,omitempty"`
	UUID           string                `json:"uuid,omitempty"`
	VendorName     string                `json:"vendor_name,omitempty"`
	VendorVersionn string                `json:"vendor_versionn,omitempty"`
	Groups         []TemplateGroupObject `json:"groups,omitempty"`
	Tags           []TemplateTag         `json:"tags,omitempty"`
	TemplatesClear []TemplateObject      `json:"templates_clear,omitempty"`
	Macros         []UserMacroObject     `json:"macros,omitempty"`
}

type TemplateResponse struct {
	TemplateIDs []string `json:"templateids,omitempty"`
}

type TemplateGet struct {
	TemplateIDs           []string      `json:"templateids,omitempty"`
	GroupIDs              []string      `json:"groupids,omitempty"`
	ParentTemplateIDs     []string      `json:"parenttemplateids,omitempty"`
	HostIDs               []string      `json:"hostids,omitempty"`
	GraphIDs              []string      `json:"graphids,omitempty"`
	ItemIDs               []string      `json:"itemids,omitempty"`
	TriggerIDs            []string      `json:"triggerids,omitempty"`
	WithItems             string        `json:"with_items,omitempty"`
	WIthTriggers          string        `json:"with_triggers,omitempty"`
	WithGraphs            string        `json:"with_graphs,omitempty"`
	WIthHttpTests         string        `json:"with_httptests,omitempty"`
	Evaltype              int           `json:"evaltype,omitempty"`
	Tags                  []TemplateTag `json:"tags,omitempty"`
	SelectTags            string        `json:"selectTags,omitempty"`
	SelectTemplateGroups  string        `json:"selectTemplateGroups,omitempty"`
	SelectTemplates       string        `json:"selectTemplates,omitempty"`
	SelectParentTemplates string        `json:"selectParentTemplates,omitempty"`
	SelectHttpTests       string        `json:"selectHttpTests,omitempty"`
	SelectItems           string        `json:"selectItems,omitempty"`
	SelectDiscoveries     string        `json:"selectDiscoveries,omitempty"`
	SelectTriggers        string        `json:"selectTriggers,omitempty"`
	SelectGraphs          string        `json:"selectGraphs,omitempty"`
	SelectMacros          string        `json:"selectMacros,omitempty"`
	SelectDashboards      string        `json:"selectDashboards,omitempty"`
	SelectValueMaps       string        `json:"selectValueMaps,omitempty"`
	LimitSelects          int           `json:"limitSelects,omitempty"`
	ZabbixCommun
}

type TemplateMassAdd struct {
	Templates     []map[string]string   `json:"templates,omitempty"`
	Groups        []TemplateGroupObject `json:"groups,omitempty"`
	Macros        []UserMacroObject     `json:"macros,omitempty"`
	TemplatesLink []map[string]string   `json:"templates_link,omitempty"`
}

type TemplateMassUpdate struct {
	Templates     []map[string]string   `json:"templates,omitempty"`
	Groups        []TemplateGroupObject `json:"groups,omitempty"`
	Macros        []UserMacroObject     `json:"macros,omitempty"`
	TemplatesLink []map[string]string   `json:"templates_link,omitempty"`
}

type TemplateMassRemove struct {
	TemplateIDs       []string          `json:"templateids,omitempty"`
	GroupIDs          []string          `json:"groupids,omitempty"`
	Macros            []UserMacroObject `json:"macros,omitempty"`
	TemplatesIDsClear []string          `json:"templateids_clear,omitempty"`
	TemplatesIDsLink  []string          `json:"templateids_link,omitempty"`
}
