package entities

type ServiceObject struct {
	ServiceID        string `json:"serviceid,omitempty"`
	Algorithm        int    `json:"algorithm,omitempty"`
	Name             string `json:"name,omitempty"`
	SortOrder        int    `json:"sortorder,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	PropagationRule  int    `json:"propagation_rule,omitempty"`
	PropagationValue int    `json:"propagation_value,omitempty"`
	Status           int    `json:"status,omitempty"`
	Description      string `json:"description,omitempty"`
	UUID             string `json:"uuid,omitempty"`
	CreatedAt        int    `json:"createdat,omitempty"`
	Readonly         bool   `json:"readonly,omitempty"`
}

type ServiceCreate struct {
	Algorithm        int                 `json:"algorithm,omitempty"`
	Name             string              `json:"name,omitempty"`
	SortOrder        int                 `json:"sortorder,omitempty"`
	Weight           int                 `json:"weight,omitempty"`
	PropagationRule  int                 `json:"propagation_rule,omitempty"`
	PropagationValue int                 `json:"propagation_value,omitempty"`
	Status           int                 `json:"status,omitempty"`
	Description      string              `json:"description,omitempty"`
	UUID             string              `json:"uuid,omitempty"`
	CreatedAt        int                 `json:"createdat,omitempty"`
	Readonly         bool                `json:"readonly,omitempty"`
	Children         []map[string]string `json:"children,omitempty"`
	Parents          []map[string]string `json:"parents,omitempty"`
	ProblemTags      []ProblemTag        `json:"problemtags,omitempty"`
	StatusRules      []ServiceStatusRule `json:"status_rules,omitempty"`
}

type ServiceUpdate struct {
	ServiceID        string              `json:"serviceid,omitempty"`
	Algorithm        int                 `json:"algorithm,omitempty"`
	Name             string              `json:"name,omitempty"`
	SortOrder        int                 `json:"sortorder,omitempty"`
	Weight           int                 `json:"weight,omitempty"`
	PropagationRule  int                 `json:"propagation_rule,omitempty"`
	PropagationValue int                 `json:"propagation_value,omitempty"`
	Status           int                 `json:"status,omitempty"`
	Description      string              `json:"description,omitempty"`
	UUID             string              `json:"uuid,omitempty"`
	CreatedAt        int                 `json:"createdat,omitempty"`
	Readonly         bool                `json:"readonly,omitempty"`
	Children         []map[string]string `json:"children,omitempty"`
	Parents          []map[string]string `json:"parents,omitempty"`
	ProblemTags      []ProblemTag        `json:"problemtags,omitempty"`
	StatusRules      []ServiceStatusRule `json:"status_rules,omitempty"`
}

type ServiceResponse struct {
	ServiceIDs []string `json:"serviceids,omitempty"`
}

type ServiceStatusRule struct {
	Type        int `json:"type,omitempty"`
	LimitValue  int `json:"limit_value,omitempty"`
	LimitStatus int `json:"limit_status,omitempty"`
	NewStatus   int `json:"new_status,omitempty"`
}

type ServiceTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

type ServiceAlarm struct {
	Clock string `json:"clock,omitempty"`
	Value int    `json:"value,omitempty"`
}

type ServiceGet struct {
	ServiceIDs           []string            `json:"serviceids,omitempty"`
	ParentIDs            []string            `json:"parentids,omitempty"`
	DeepParentIDs        string              `json:"deep_parentids,omitempty"`
	ChildIDs             []string            `json:"childids,omitempty"`
	Evaltype             int                 `json:"evaltype,omitempty"`
	Tags                 []ZabbixTag         `json:"tags,omitempty"`
	ProblemTags          []ProblemTag        `json:"problem_tags,omitempty"`
	WithoutProblemTags   string              `json:"without_problem_tags,omitempty"`
	Slaids               []string            `json:"slaids,omitempty"`
	SelectChildren       string              `json:"selectChildren,omitempty"`
	SelectParents        string              `json:"selectParents,omitempty"`
	SelectTags           string              `json:"selectTags,omitempty"`
	SelectProblemEvents  string              `json:"selectProblemEvents,omitempty"`
	SelectProblemTags    string              `json:"selectProblemTags,omitempty"`
	SelectStatusRules    string              `json:"selectStatusRules,omitempty"`
	SelectStatusTimeline []map[string]string `json:"selectStatustimeLine,omitempty"`
	ZabbixCommun
}
