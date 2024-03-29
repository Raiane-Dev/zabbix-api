package entities

type CorrelationCreate struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Status      int                    `json:"status"`
	Operations  []CorrelationOperation `json:"operations"`
	Filter      CorrelationFilter      `json:"filter"`
}

type CorrelationUpdate struct {
	CorrelationID string
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	Status        int                    `json:"status"`
	Operations    []CorrelationOperation `json:"operations"`
	Filter        CorrelationFilter      `json:"filter"`
}

type CorrelationResponse struct {
	CorrelationIDs []string `json:"correlationids"`
}

type CorrelationGet struct {
	CorrelationIDs   []string `json:"correlationid"`
	SelectFilter     string   `json:"selectFilter"`
	SelectOperations string   `json:"selectOperations"`
	ZabbixCommun
}

type CorrelationOutput struct {
	CorrelationID string
	Name          string
	Description   string
	Status        string
	Filter        CorrelationFilter
}

type CorrelationOperation struct {
	Type int `json:"type"`
}

type CorrelationFilter struct {
	Evaltype    int                          `json:"evaltype"`
	Conditions  []CorrelationFilterCondition `json:"conditions"`
	EvalFormula string                       `json:"eval_formula"`
	Formula     string                       `json:"formula"`
}

type CorrelationFilterCondition struct {
	Type      int    `json:"type"`
	Tag       string `json:"tag"`
	GroupID   string `json:"groupid"`
	OldTag    string `json:"oldtag"`
	NewTag    string `json:"newtag"`
	Value     string `json:"value"`
	FormulaID string `json:"formulaid"`
	Operator  int    `json:"operator"`
}
