package entities

type ConfigurationExport struct {
	Format      string              `json:"format"`
	PrettyPrint bool                `json:"prettyprint"`
	Options     map[string][]string `json:"options"`
}

type ConfigurationImport struct {
	Format string
	Source string
	Rules  ConfigurationRules
}

type ConfigurationRules struct {
	DiscoveryRules     map[string]string `json:"discoveryRules"`
	Graphs             map[string]string `json:"graphs"`
	HostGroups         map[string]string `json:"host_groups"`
	TemplateGroups     map[string]string `json:"template_groups"`
	Hosts              map[string]string `json:"hosts"`
	HttpTests          map[string]string `json:"httptests"`
	Images             map[string]string `json:"images"`
	Items              map[string]string `json:"items"`
	Maps               map[string]string `json:"maps"`
	MediaTypes         map[string]string `json:"mediaTypes"`
	TemplateLinkage    map[string]string `json:"templateLinkage"`
	Templates          map[string]string `json:"templates"`
	TemplateDashboards map[string]string `json:"templateDashboards"`
	Triggers           map[string]string `json:"triggers"`
	ValueMaps          map[string]string `json:"valueMaps"`
}

type ConfigurationComparedOutput struct {
	Templates ConfigurationCompareUpdatedOutput `jaon:"templates"`
}
type ConfigurationCompareUpdatedOutput struct {
	Updated []ConfigurationCompareUpdatedOutput `json:"updated"`
}

type ConfigurationCompareTriggerOutput struct {
	UUID       string `json:"uuid"`
	Expression string `json:"expression"`
	Name       string `json:"name"`
	Priority   string `json:"priority"`
}

type ConfigurationCompareItemOutput struct {
	UUID      string                             `json:"uuid"`
	Name      string                             `json:"name"`
	Key       string                             `json:"key"`
	ValueType string                             `json:"value_type"`
	Units     string                             `json:"units"`
	Triggers  ConfigurationCompareTriggersOutput `json:"triggers"`
}

type ConfigurationCompareTriggersOutput struct {
	Added   []ConfigurationCompareTriggerOutput `json:"added"`
	Removed []ConfigurationCompareTriggerOutput `json:"removed"`
}

type ConfigurationCompareTemplateUpdateOutput struct {
	Before ConfigurationCompareTemplateOutput `json:"before"`
	After  ConfigurationCompareTemplateOutput `json:"after"`
	Items  ItemUpdate                         `json:"items"`
}

type ConfigurationCompareTemplateOutput struct {
	UUID   string                            `json:"uuid"`
	Name   string                            `json:"name"`
	Groups []ConfigurationCompareGroupOutput `json:"groups"`
}

type ConfigurationCompareGroupOutput struct {
	Name string `json:"name"`
}

type ConfigurationCompareItemUpdateOutput struct {
	Added   []ConfigurationCompareItemOutput `json:"added"`
	Removed []ConfigurationCompareItemOutput `json:"removed"`
	Updated []ConfigurationCompareItemOutput `json:"updated"`
}
