package entities

import "time"

type TriggerObject struct {
	TriggerID          string    `json:"triggerid,omitempty"`
	Description        string    `json:"description,omitempty"`
	Expression         string    `json:"expression,omitempty"`
	EventName          string    `json:"event_name,omitempty"`
	OpData             string    `json:"opdata,omitempty"`
	Comments           string    `json:"comments,omitempty"`
	Error              string    `json:"error,omitempty"`
	Flags              int       `json:"flags,omitempty"`
	LastChange         time.Time `json:"lastchange,omitempty"`
	Priority           int       `json:"priority,omitempty"`
	State              int       `json:"state,omitempty"`
	Status             int       `json:"status,omitempty"`
	TemplateID         string    `json:"templateid,omitempty"`
	Type               int       `json:"type,omitempty"`
	URL                string    `json:"url,omitempty"`
	URLName            string    `json:"url_name,omitempty"`
	Value              int       `json:"value,omitempty"`
	RecoveryMode       int       `json:"recovery_mode,omitempty"`
	RecoveryExpression string    `json:"recovery_expression,omitempty"`
	CorrelationMode    int       `json:"correlation_mode,omitempty"`
	CorrelationTag     string    `json:"correlation_tag,omitempty"`
	ManualClose        int       `json:"manual_close,omitempty"`
	UUID               string    `json:"uuid,omitempty"`
}

type TriggerCreate struct {
	Description        string              `json:"description,omitempty"`
	Expression         string              `json:"expression,omitempty"`
	EventName          string              `json:"event_name,omitempty"`
	OpData             string              `json:"opdata,omitempty"`
	Comments           string              `json:"comments,omitempty"`
	Error              string              `json:"error,omitempty"`
	Flags              int                 `json:"flags,omitempty"`
	LastChange         time.Time           `json:"lastchange,omitempty"`
	Priority           int                 `json:"priority,omitempty"`
	State              int                 `json:"state,omitempty"`
	Status             int                 `json:"status,omitempty"`
	TemplateID         string              `json:"templateid,omitempty"`
	Type               int                 `json:"type,omitempty"`
	URL                string              `json:"url,omitempty"`
	URLName            string              `json:"url_name,omitempty"`
	Value              int                 `json:"value,omitempty"`
	RecoveryMode       int                 `json:"recovery_mode,omitempty"`
	RecoveryExpression string              `json:"recovery_expression,omitempty"`
	CorrelationMode    int                 `json:"correlation_mode,omitempty"`
	CorrelationTag     string              `json:"correlation_tag,omitempty"`
	ManualClose        int                 `json:"manual_close,omitempty"`
	UUID               string              `json:"uuid,omitempty"`
	Dependencies       []map[string]string `json:"dependencies,omitempty"`
	Tags               []ZabbixTag         `json:"tags,omitempty"`
}

type TriggerUpdate struct {
	TriggerID          string              `json:"triggerid,omitempty"`
	Description        string              `json:"description,omitempty"`
	Expression         string              `json:"expression,omitempty"`
	EventName          string              `json:"event_name,omitempty"`
	OpData             string              `json:"opdata,omitempty"`
	Comments           string              `json:"comments,omitempty"`
	Error              string              `json:"error,omitempty"`
	Flags              int                 `json:"flags,omitempty"`
	LastChange         time.Time           `json:"lastchange,omitempty"`
	Priority           int                 `json:"priority,omitempty"`
	State              int                 `json:"state,omitempty"`
	Status             int                 `json:"status,omitempty"`
	TemplateID         string              `json:"templateid,omitempty"`
	Type               int                 `json:"type,omitempty"`
	URL                string              `json:"url,omitempty"`
	URLName            string              `json:"url_name,omitempty"`
	Value              int                 `json:"value,omitempty"`
	RecoveryMode       int                 `json:"recovery_mode,omitempty"`
	RecoveryExpression string              `json:"recovery_expression,omitempty"`
	CorrelationMode    int                 `json:"correlation_mode,omitempty"`
	CorrelationTag     string              `json:"correlation_tag,omitempty"`
	ManualClose        int                 `json:"manual_close,omitempty"`
	UUID               string              `json:"uuid,omitempty"`
	Dependencies       []map[string]string `json:"dependencies,omitempty"`
	Tags               []ZabbixTag         `json:"tags,omitempty"`
}

type TriggerResponse struct {
	TriggerIDs []string `json:"triggerids"`
}

type TriggerGet struct {
	TriggerIDs                  []string    `json:"triggerids,omitempty"`
	GroupIds                    []string    `json:"groupids,omitempty"`
	TemplateIds                 []string    `json:"templateids,omitempty"`
	HostIds                     []string    `json:"hostids,omitempty"`
	ItemIds                     []string    `json:"itemids,omitempty"`
	Functions                   []string    `json:"functions,omitempty"`
	Group                       string      `json:"group,omitempty"`
	Host                        string      `json:"host,omitempty"`
	Inherited                   bool        `json:"inherited,omitempty"`
	Templated                   bool        `json:"templated,omitempty"`
	Dependent                   bool        `json:"dependent,omitempty"`
	Monitored                   bool        `json:"monitored,omitempty"`
	Active                      bool        `json:"active,omitempty"`
	Maintenance                 bool        `json:"maintenance,omitempty"`
	WithUnacknowledgedEvents    bool        `json:"withUnacknowledgedEvents,omitempty"`
	WithAcknowledgedEvents      bool        `json:"withAcknowledgedEvents,omitempty"`
	WithLastEventUnacknowledged bool        `json:"withLastEventUnacknowledged,omitempty"`
	SkipDependent               bool        `json:"skipDependent,omitempty"`
	LastChangeSince             string      `json:"lastChangeSince,omitempty"`
	LastChangeTill              string      `json:"lastChangeTill,omitempty"`
	OnlyTrue                    bool        `json:"only_true,omitempty"`
	MinSeverity                 int         `json:"min_severity,omitempty"`
	EvalType                    int         `json:"evaltype,omitempty"`
	Tags                        []TagFilter `json:"tags,omitempty"`
	ExpandComment               bool        `json:"expandComment,omitempty"`
	ExpandDescription           bool        `json:"expandDescription,omitempty"`
	ExpandExpression            bool        `json:"expandExpression,omitempty"`
	SelectHostGroups            bool        `json:"selectHostGroups,omitempty"`
	SelectHosts                 bool        `json:"selectHosts,omitempty"`
	SelectItems                 bool        `json:"selectItems,omitempty"`
	SelectFunctions             bool        `json:"selectFunctions,omitempty"`
	SelectDependencies          bool        `json:"selectDependencies,omitempty"`
	SelectDiscoveryRule         bool        `json:"selectDiscoveryRule,omitempty"`
	SelectLastEvent             bool        `json:"selectLastEvent,omitempty"`
	SelectTags                  bool        `json:"selectTags,omitempty"`
	SelectTemplateGroups        bool        `json:"selectTemplateGroups,omitempty"`
	SelectTriggerDiscovery      bool        `json:"selectTriggerDiscovery,omitempty"`
	ZabbixCommun
}

type TagFilter struct {
	Tag      string `json:"tag,omitempty"`
	Value    string `json:"value,omitempty"`
	Operator string `json:"operator,omitempty"`
}
