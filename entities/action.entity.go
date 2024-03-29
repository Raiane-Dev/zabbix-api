package entities

type ActionGet struct {
	ActionIDs                []string `json:"actionids,omitempty"`
	GroupIDs                 []string `json:"groupids,omitempty"`
	HostIDs                  []string `json:"hostids,omitempty"`
	TriggerIDs               []string `json:"triggerids,omitempty"`
	UsrGrpIDs                []string `json:"usrgrpids,omitempty"`
	UserIDs                  []string `json:"userids,omitempty"`
	ScriptIDs                []string `json:"scriptids,omitempty"`
	SelectFilter             string   `json:"selectFilter,omitempty"`
	SelectOperations         string   `json:"selectOperations,omitempty"`
	SelectRecoveryOperations string   `json:"selectRecoveryOperations,omitempty"`
	SelectUpdateOperations   string   `json:"selectUpdateOperations,omitempty"`
	ZabbixCommun
}

type ActionObject struct {
	ActionID           string                    `json:"actionid,omitempty"`
	Name               string                    `json:"name,omitempty"`
	EventSource        string                    `json:"eventsource,omitempty"`
	EscPeriod          string                    `json:"esc_period,omitempty"`
	PauseSupressed     string                    `json:"pause_supressed,omitempty"`
	Filter             ActionFilter              `json:"filter,omitempty"`
	Operations         []ActionOperationsObject  `json:"operations,omitempty"`
	RecoveryOperations []RecoveryOperationObject `json:"recovery_operations,omitempty"`
}

type RecoveryOperationObject struct {
	OperationID   string                     `json:"operationid,omitempty"`
	ActionID      string                     `json:"actionid,omitempty"`
	OperationType string                     `json:"operationtype,omitempty"`
	Evaltype      string                     `json:"evaltype,omitempty"`
	OpConditions  []ActionOperationCondition `json:"opconditions,omitempty"`
	OpMessage     ActionOperationMessage     `json:"opmessage,omitempty"`
}

type ActionOperationsObject struct {
	OperationID   string                   `json:"operationid,omitempty"`
	ActionID      string                   `json:"actionid,omitempty"`
	OperationType int                      `json:"operationtype,omitempty"`
	EscPeriod     int                      `json:"esc_period,omitempty"`
	EscStepFrom   int                      `json:"esc_step_from,omitempty"`
	EscStepTo     int                      `json:"esc_step_to,omitempty"`
	Evaltype      string                   `json:"evaltype,omitempty"`
	OpConditions  []string                 `json:"opconditions,omitempty"`
	OpMessage     []ActionOperationMessage `json:"opmessage,omitempty"`
	OpMessageGRP  []map[string]string      `json:"opmessage_grp,omitempty"`
}

type ActionCreate struct {
	ActionID         string `json:"actionid,omitempty"`
	EscPeriod        string `json:"esc_period,omitempty"`
	EventSource      int    `json:"eventsource,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           int    `json:"status,omitempty"`
	PauseSymptoms    int    `json:"pause_symptoms,omitempty"`
	PauseSupressed   int    `json:"pause_supressed,omitempty"`
	NotifyIfCanceled int    `json:"notify_if_canceled,omitempty"`
}

type ActionOperation struct {
	OperationType int                        `json:"operationtype,omitempty"`
	EscPeriod     string                     `json:"esc_period,omitempty"`
	EscStepFrom   int                        `json:"esc_step_from,omitempty"`
	EscStepTo     int                        `json:"esc_step_to,omitempty"`
	Evaltype      int                        `json:"evaltype,omitempty"`
	OpCommand     map[string]string          `json:"opcommand,omitempty"`
	OpCommandGRP  []map[string]string        `json:"opcommand_grp,omitempty"`
	OpCommandHST  []map[string]string        `json:"opcommand_hst,omitempty"`
	OpConditions  []ActionOperationCondition `json:"opconditions,omitempty"`
	OpGroup       []string                   `json:"opgroup,omitempty"`
	OpMessage     ActionOperationMessage     `json:"opmessage,omitempty"`
	OpMessageGRP  []string                   `json:"opmessage_grp,omitempty"`
	OpMessageUSR  []string                   `json:"opmessage_usr,omitempty"`
	OpTemplate    interface{}                `json:"optemplate,omitempty"`
	OpInventory   map[string]string          `json:"opinventory,omitempty"`
}

type ActionOperationMessage struct {
	DefaultMsg  int    `json:"default_msg,omitempty"`
	MediaTypeID string `json:"mediatypeid,omitempty"`
	Message     string `json:"message,omitempty"`
	Subject     string `json:"subject,omitempty"`
}

type ActionOperationCondition struct {
	ConditionType int    `json:"conditiontype,omitempty"`
	Value         string `json:"value,omitempty"`
	Operator      int    `json:"operator,omitempty"`
}

type ActionRecoveryOperation struct {
	OperationType int                    `json:"operationtype,omitempty"`
	OpCommand     map[string]string      `json:"opcommand,omitempty"`
	OpCommandGRP  []map[string]string    `json:"opcommand_grp,omitempty"`
	OpCommandHST  []map[string]string    `json:"opcommand_hst,omitempty"`
	OpMessage     ActionOperationMessage `json:"opmessage,omitempty"`
	OpMessageGRP  []map[string]string    `json:"opmessage_grp,omitempty"`
	OpMessageUSR  []map[string]string    `json:"opmessage_usr,omitempty"`
}

type ActionUpdate struct {
	OperationType int
	OpCommand     map[string]string      `json:"opcommand,omitempty"`
	OpCommandGRP  []string               `json:"opcommand_grp,omitempty"`
	OpCommandHST  []string               `json:"opcommand_hst,omitempty"`
	OpMessage     ActionOperationMessage `json:"opmessage,omitempty"`
	OpMessageGRP  []map[string]string    `json:"opmessage_grp,omitempty"`
	OpMessageUSR  []map[string]string    `json:"opmessage_usr,omitempty"`
}

type ActionFilter struct {
	Conditions  []string `json:"conditions,omitempty"`
	Evaltype    int      `json:"evaltype,omitempty"`
	EvalFormula string   `json:"eval_formula,omitempty"`
	Formula     string   `json:"formula,omitempty"`
}

type ActionFilterCondition struct {
	ConditionType int    `json:"conditiontype,omitempty"`
	Value         string `json:"value,omitempty"`
	Value2        string `json:"value2,omitempty"`
	FormulaID     string `json:"formulaid,omitempty"`
	Operator      int    `json:"operator,omitempty"`
}

type ActionResponse struct {
	ActionIDs []string `json:"actionids"`
}
