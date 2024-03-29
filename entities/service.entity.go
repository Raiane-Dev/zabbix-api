package entities

type ServiceObject struct {
	ServiceID        string `json:"serviceid"`
	Algorithm        int    `json:"algorithm"`
	Name             string `json:"name"`
	SortOrder        int    `json:"sortorder"`
	Weight           int    `json:"weight"`
	PropagationRule  int    `json:"propagation_rule"`
	PropagationValue int    `json:"propagation_value"`
	Status           int    `json:"status"`
	Description      string `json:"description"`
	UUID             string `json:"uuid"`
	CreatedAt        int    `json:"createdat"`
	Readonly         bool   `json:"readonly"`
}

type ServiceCreate struct {
	Algorithm        int                 `json:"algorithm"`
	Name             string              `json:"name"`
	SortOrder        int                 `json:"sortorder"`
	Weight           int                 `json:"weight"`
	PropagationRule  int                 `json:"propagation_rule"`
	PropagationValue int                 `json:"propagation_value"`
	Status           int                 `json:"status"`
	Description      string              `json:"description"`
	UUID             string              `json:"uuid"`
	CreatedAt        int                 `json:"createdat"`
	Readonly         bool                `json:"readonly"`
	Children         []map[string]string `json:"children"`
	Parents          []map[string]string `json:"parents"`
	ProblemTags      []ProblemTag        `json:"problemtags"`
	StatusRules      []ServiceStatusRule `json:"statusrules"`
}

type ServiceStatusRule struct {
	Type        int `json:"type"`
	LimitValue  int `json:"limit_value"`
	LimitStatus int `json:"limit_status"`
	NewStatus   int `json:"new_status"`
}

type ServiceTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ServiceAlarm struct {
	Clock string `json:"clock"`
	Value int    `json:"value"`
}
