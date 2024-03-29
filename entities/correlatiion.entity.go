package entities

type CorrelationCreate struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Status      int                    `json:"status,omitempty"`
	Operations  []CorrelationOperation `json:"operations,omitempty"`
	Filter      CorrelationFilter      `json:"filter,omitempty"`
}

type CorrelationUpdate struct {
	CorrelationID string
	Name          string                 `json:"name,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Status        int                    `json:"status,omitempty"`
	Operations    []CorrelationOperation `json:"operations,omitempty"`
	Filter        CorrelationFilter      `json:"filter,omitempty"`
}

type CorrelationResponse struct {
	CorrelationIDs []string `json:"correlationids,omitempty"`
}

type CorrelationGet struct {
	CorrelationIDs   []string `json:"correlationid,omitempty"`
	SelectFilter     string   `json:"selectFilter,omitempty"`
	SelectOperations string   `json:"selectOperations,omitempty"`
	ZabbixCommun
}

type CorrelationObject struct {
	CorrelationID string            `json:"correlationid,omitempty"`
	Name          string            `json:"name,omitempty"`
	Description   string            `json:"description,omitempty"`
	Status        string            `json:"status,omitempty"`
	Filter        CorrelationFilter `json:"filter,omitempty"`
}

type CorrelationOperation struct {
	Type int `json:"type,omitempty"`
}

type CorrelationFilter struct {
	Evaltype    int                          `json:"evaltype,omitempty"`
	Conditions  []CorrelationFilterCondition `json:"conditions,omitempty"`
	EvalFormula string                       `json:"eval_formula,omitempty"`
	Formula     string                       `json:"formula,omitempty"`
}

type CorrelationFilterCondition struct {
	Type      int    `json:"type,omitempty"`
	Tag       string `json:"tag,omitempty"`
	GroupID   string `json:"groupid,omitempty"`
	OldTag    string `json:"oldtag,omitempty"`
	NewTag    string `json:"newtag,omitempty"`
	Value     string `json:"value,omitempty"`
	FormulaID string `json:"formulaid,omitempty"`
	Operator  int    `json:"operator,omitempty"`
}
