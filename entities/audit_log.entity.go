package entities

type AuditlogGet struct {
	AuditIDs []string `json:"auditids,omitempty"`
	UserIDs  []string `json:"userids,omitempty"`
	TimeFrom string   `json:"time_from,omitempty"`
	TimeTill string   `json:"time_till,omitempty"`
	ZabbixCommun
}

type AuditlogObject struct {
	AuditID      string `json:"auditid,omitempty"`
	UserID       string `json:"userid,omitempty"`
	Username     string `json:"username,omitempty"`
	Clock        string `json:"clock,omitempty"`
	IP           string `json:"ip,omitempty"`
	Action       string `json:"action,omitempty"`
	ResourceType string `json:"resourcetype,omitempty"`
	ResourceID   string `json:"resourceid,omitempty"`
	ResourceName string `json:"resourcename,omitempty"`
	RecordsetID  string `json:"recordsetid,omitempty"`
	Details      string `json:"details,omitempty"`
}
