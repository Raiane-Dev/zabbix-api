package entities

type TemplateGroupObject struct {
	GroupID string `json:"groupid,omitempty"`
	Name    string `json:"name,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type TemplateGroupCreate struct {
	Name string `json:"name,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type TemplateGroupUpdate struct {
	GroupID string `json:"groupid,omitempty"`
	Name    string `json:"name,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}
type TemplateGroupResponse struct {
	GroupIDs []string `json:"groupids,omitempty"`
}

type TemplateGroupGet struct {
	GraphIds                      []string `json:"graphids,omitempty"`
	GroupIds                      []string `json:"groupids,omitempty"`
	TemplateIds                   []string `json:"templateids,omitempty"`
	TriggerIds                    []string `json:"triggerids,omitempty"`
	WithGraphs                    bool     `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool     `json:"with_graph_prototypes,omitempty"`
	WithHttpTests                 bool     `json:"with_httptests,omitempty"`
	WithItems                     bool     `json:"with_items,omitempty"`
	WithItemPrototypes            bool     `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool     `json:"with_simple_graph_item_prototypes,omitempty"`
	WithSimpleGraphItems          bool     `json:"with_simple_graph_items,omitempty"`
	WithTemplates                 bool     `json:"with_templates,omitempty"`
	WithTriggers                  bool     `json:"with_triggers,omitempty"`
	SelectTemplates               string   `json:"selectTemplates,omitempty"`
	LimitSelects                  int      `json:"limitSelects,omitempty"`
	ZabbixCommun
}

type TemplateGroupMassAdd struct {
	Groups    []map[string]string `json:"groups,omitempty"`
	Templates []map[string]string `json:"templates,omitempty"`
}

type TemplateGroupMassRemove struct {
	Groups    []map[string]string `json:"groups,omitempty"`
	Templates []map[string]string `json:"templates,omitempty"`
}

type TemplateGroupMassUpdate struct {
	Groups    []map[string]string `json:"groups,omitempty"`
	Templates []map[string]string `json:"templates,omitempty"`
}

type TemplateGroupPropagate struct {
	Groups      []map[string]string `json:"groups,omitempty"`
	Permissions bool                `json:"permissions,omitempty"`
}
