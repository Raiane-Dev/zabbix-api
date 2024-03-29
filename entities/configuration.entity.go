package entities

type ConfigurationExport struct {
	Format      string              `json:"format,omitempty"`
	PrettyPrint bool                `json:"prettyprint,omitempty"`
	Options     map[string][]string `json:"options,omitempty"`
}

type ConfigurationImport struct {
	Format string                     `json:"format,omitempty"`
	Source string                     `json:"source,omitempty"`
	Rules  ConfigurationImportCompare `json:"rules,omitempty"`
}

type ConfigurationImportCompare struct {
	DiscoveryRules     map[string]any `json:"discoveryRules,omitempty"`
	Graphs             map[string]any `json:"graphs,omitempty"`
	HostGroups         map[string]any `json:"host_groups,omitempty"`
	TemplateGroups     map[string]any `json:"template_groups,omitempty"`
	Hosts              map[string]any `json:"hosts,omitempty"`
	HttpTests          map[string]any `json:"httptests,omitempty"`
	Images             map[string]any `json:"images,omitempty"`
	Items              map[string]any `json:"items,omitempty"`
	Maps               map[string]any `json:"maps,omitempty"`
	MediaTypes         map[string]any `json:"mediaTypes,omitempty"`
	TemplateLinkage    map[string]any `json:"templateLinkage,omitempty"`
	Templates          map[string]any `json:"templates,omitempty"`
	TemplateDashboards map[string]any `json:"templateDashboards,omitempty"`
	Triggers           map[string]any `json:"triggers,omitempty"`
	ValueMaps          map[string]any `json:"valueMaps,omitempty"`
}

type ConfigurationCompareResponse struct {
	Templates ConfigurationCompareUpdatedObject `jaon:"templates,omitempty"`
}
type ConfigurationCompareUpdatedObject struct {
	Updated []ConfigurationCompareUpdatedObject `json:"updated,omitempty"`
}

type ConfigurationCompareTriggerObject struct {
	UUID       string `json:"uuid,omitempty"`
	Expression string `json:"expression,omitempty"`
	Name       string `json:"name,omitempty"`
	Priority   string `json:"priority,omitempty"`
}

type ConfigurationCompareItemObject struct {
	UUID      string                             `json:"uuid,omitempty"`
	Name      string                             `json:"name,omitempty"`
	Key       string                             `json:"key,omitempty"`
	ValueType string                             `json:"value_type,omitempty"`
	Units     string                             `json:"units,omitempty"`
	Triggers  ConfigurationCompareTriggersObject `json:"triggers,omitempty"`
}

type ConfigurationCompareTriggersObject struct {
	Added   []ConfigurationCompareTriggerObject `json:"added,omitempty"`
	Removed []ConfigurationCompareTriggerObject `json:"removed,omitempty"`
}

type ConfigurationCompareTemplateUpdateObject struct {
	Before ConfigurationCompareTemplateObject `json:"before,omitempty"`
	After  ConfigurationCompareTemplateObject `json:"after,omitempty"`
	Items  ItemUpdate                         `json:"items,omitempty"`
}

type ConfigurationCompareTemplateObject struct {
	UUID   string                            `json:"uuid,omitempty"`
	Name   string                            `json:"name,omitempty"`
	Groups []ConfigurationCompareGroupObject `json:"groups,omitempty"`
}

type ConfigurationCompareGroupObject struct {
	Name string `json:"name,omitempty"`
}

type ConfigurationCompareItemUpdateObject struct {
	Added   []ConfigurationCompareItemObject `json:"added,omitempty"`
	Removed []ConfigurationCompareItemObject `json:"removed,omitempty"`
	Updated []ConfigurationCompareItemObject `json:"updated,omitempty"`
}
