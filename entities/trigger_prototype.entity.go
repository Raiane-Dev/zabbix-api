package entities

type TriggerPrototypeObject struct {
	TriggerID          string `json:"triggerid,omitempty"`
	Description        string `json:"description,omitempty"`
	Expression         string `json:"expression,omitempty"`
	EventName          string `json:"event_name,omitempty"`
	OpData             string `json:"opdata,omitempty"`
	Comments           string `json:"comments,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	Status             int    `json:"status,omitempty"`
	TemplateID         string `json:"templateid,omitempty"`
	Type               int    `json:"type,omitempty"`
	URL                string `json:"url,omitempty"`
	URLName            string `json:"url_name,omitempty"`
	RecoveryMode       int    `json:"recovery_mode,omitempty"`
	RecoveryExpression string `json:"recovery_expression,omitempty"`
	CorrelationMode    int    `json:"correlation_mode,omitempty"`
	CorrelationTag     string `json:"correlation_tag,omitempty"`
	ManualClose        int    `json:"manual_close,omitempty"`
	Discover           int    `json:"discover,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}

type TriggerPrototypeCreate struct {
	Description        string              `json:"description,omitempty"`
	Expression         string              `json:"expression,omitempty"`
	EventName          string              `json:"event_name,omitempty"`
	OpData             string              `json:"opdata,omitempty"`
	Comments           string              `json:"comments,omitempty"`
	Priority           int                 `json:"priority,omitempty"`
	Status             int                 `json:"status,omitempty"`
	TemplateID         string              `json:"templateid,omitempty"`
	Type               int                 `json:"type,omitempty"`
	URL                string              `json:"url,omitempty"`
	URLName            string              `json:"url_name,omitempty"`
	RecoveryMode       int                 `json:"recovery_mode,omitempty"`
	RecoveryExpression string              `json:"recovery_expression,omitempty"`
	CorrelationMode    int                 `json:"correlation_mode,omitempty"`
	CorrelationTag     string              `json:"correlation_tag,omitempty"`
	ManualClose        int                 `json:"manual_close,omitempty"`
	Discover           int                 `json:"discover,omitempty"`
	UUID               string              `json:"uuid,omitempty"`
	Dependencies       []map[string]string `json:"dependencies,omitempty"`
	Tags               []ZabbixTag         `json:"tags,omitempty"`
}

type TriggerPrototypeUpdate struct {
	TriggerID          string              `json:"triggerid,omitempty"`
	Description        string              `json:"description,omitempty"`
	Expression         string              `json:"expression,omitempty"`
	EventName          string              `json:"event_name,omitempty"`
	OpData             string              `json:"opdata,omitempty"`
	Comments           string              `json:"comments,omitempty"`
	Priority           int                 `json:"priority,omitempty"`
	Status             int                 `json:"status,omitempty"`
	TemplateID         string              `json:"templateid,omitempty"`
	Type               int                 `json:"type,omitempty"`
	URL                string              `json:"url,omitempty"`
	URLName            string              `json:"url_name,omitempty"`
	RecoveryMode       int                 `json:"recovery_mode,omitempty"`
	RecoveryExpression string              `json:"recovery_expression,omitempty"`
	CorrelationMode    int                 `json:"correlation_mode,omitempty"`
	CorrelationTag     string              `json:"correlation_tag,omitempty"`
	ManualClose        int                 `json:"manual_close,omitempty"`
	Discover           int                 `json:"discover,omitempty"`
	UUID               string              `json:"uuid,omitempty"`
	Dependencies       []map[string]string `json:"dependencies,omitempty"`
	Tags               []ZabbixTag         `json:"tags,omitempty"`
}

type TriggerPrototypeResponse struct {
	TriggerIDs []string `json:"triggerids"`
}

type TriggerPrototypeGet struct {
	Active               bool     `json:"active,omitempty"`
	DiscoveryIDs         []string `json:"discoveryids,omitempty"`
	Functions            []string `json:"functions,omitempty"`
	Group                string   `json:"group,omitempty"`
	GroupIDs             []string `json:"groupids,omitempty"`
	Host                 string   `json:"host,omitempty"`
	HostIDs              []string `json:"hostids,omitempty"`
	Inherited            bool     `json:"inherited,omitempty"`
	Maintenance          bool     `json:"maintenance,omitempty"`
	MinSeverity          int      `json:"min_severity,omitempty"`
	Monitored            bool     `json:"monitored,omitempty"`
	Templated            bool     `json:"templated,omitempty"`
	TemplateIDs          []string `json:"templateids,omitempty"`
	TriggerIDs           []string `json:"triggerids,omitempty"`
	ExpandExpression     bool     `json:"expandExpression,omitempty"`
	SelectDependencies   bool     `json:"selectDependencies,omitempty"`
	SelectDiscoveryRule  bool     `json:"selectDiscoveryRule,omitempty"`
	SelectFunctions      bool     `json:"selectFunctions,omitempty"`
	SelectHostGroups     bool     `json:"selectHostGroups,omitempty"`
	SelectHosts          bool     `json:"selectHosts,omitempty"`
	SelectItems          bool     `json:"selectItems,omitempty"`
	SelectTags           bool     `json:"selectTags,omitempty"`
	SelectTemplateGroups bool     `json:"selectTemplateGroups,omitempty"`
}
