package entities

type EventGet struct {
	EventIDs             []string            `json:"eventids,omitempty"`
	GroupIDs             []string            `json:"groupids,omitempty"`
	HostIDs              []string            `json:"hostids,omitempty"`
	Source               int                 `json:"source,omitempty"`
	Object               int                 `json:"object,omitempty"`
	Acknowledged         bool                `json:"acknowledged,omitempty"`
	Supressed            bool                `json:"supressed,omitempty"`
	Symptom              bool                `json:"symptom,omitempty"`
	Severities           []int               `json:"severities,omitempty"`
	Evaltype             int                 `json:"evaltype,omitempty"`
	Tags                 []map[string]string `json:"tags,omitempty"`
	EventIDFrom          string              `json:"eventid_from,omitempty"`
	EventIDTill          string              `json:"eventid_till,omitempty"`
	TimeFrom             string              `json:"time_from,omitempty"`
	TimeTill             string              `json:"time_till,omitempty"`
	ProblemTimeFrom      string              `json:"problem_time_from,omitempty"`
	ProblemTimeTill      string              `json:"problem_time_till,omitempty"`
	Value                []int               `json:"value,omitempty"`
	SelectHosts          string              `json:"selectHosts,omitempty"`
	SelectRelatedObject  string              `json:"selectRelatedObject,omitempty"`
	SelectAlerts         string              `json:"select_alerts,omitempty"`
	SelectAcknowledges   string              `json:"select_acknowledges,omitempty"`
	SelectTags           string              `json:"selectTags,omitempty"`
	SelectSupressionData string              `json:"selectSupressionData,omitempty"`
	ZabbixCommun
}

type EventObject struct {
	EventID        string                `json:"eventid,omitempty"`
	Source         string                `json:"source,omitempty"`
	Object         string                `json:"object,omitempty"`
	ObjectID       string                `json:"objectid,omitempty"`
	Clock          string                `json:"clock,omitempty"`
	Value          string                `json:"value,omitempty"`
	Acknowledged   string                `json:"acknowledged,omitempty"`
	NS             string                `json:"ns,omitempty"`
	Name           string                `json:"name,omitempty"`
	Severity       string                `json:"severity,omitempty"`
	REventID       string                `json:"r_eventid,omitempty"`
	CEventID       string                `json:"c_eventid,omitempty"`
	CorrelationID  string                `json:"correlationid,omitempty"`
	UserID         string                `json:"userid,omitempty"`
	CauseEventID   string                `json:"cause_eventid,omitempty"`
	Opdata         string                `json:"opdata,omitempty"`
	Acknowledges   []EventAcknowledge    `json:"acknowledges,omitempty"`
	SupressionData []EventSupressionData `json:"supressiondata,omitempty"`
	Supressed      string                `json:"supressed,omitempty"`
	Tags           []map[string]string   `json:"tags,omitempty"`
}

type EventAcknowledge struct {
	AcknowledgeID string `json:"acknowledgeid,omitempty"`
	UserID        string `json:"userid,omitempty"`
	EventID       string `json:"eventid,omitempty"`
	Clock         string `json:"clock,omitempty"`
	Message       string `json:"message,omitempty"`
	Action        string `json:"action,omitempty"`
	OldSeverity   string `json:"old_severity,omitempty"`
	NewSeverity   string `json:"new_severity,omitempty"`
	SupressUntil  string `json:"supress_until,omitempty"`
	TaskID        string `json:"taskid,omitempty"`
	Username      string `json:"username,omitempty"`
	Name          string `json:"name,omitempty"`
	Surname       string `json:"surname,omitempty"`
}

type EventResponse struct {
	EventIDs []string `json:"eventids,omitempty"`
}

type EventSupressionData struct {
	MaintenanceID string `json:"maintenanceid,omitempty"`
	SuppressUntil string `json:"suppress_until,omitempty"`
	UserID        string `json:"userid,omitempty"`
}
