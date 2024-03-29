package entities

type EventGet struct {
	EventIDs             []string            `json:"eventids"`
	GroupIDs             []string            `json:"groupids"`
	HostIDs              []string            `json:"hostids"`
	Source               int                 `json:"source"`
	Object               int                 `json:"object"`
	Acknowledged         bool                `json:"acknowledged"`
	Supressed            bool                `json:"supressed"`
	Symptom              bool                `json:"symptom"`
	Severities           []int               `json:"severities"`
	Evaltype             int                 `json:"evaltype"`
	Tags                 []map[string]string `json:"tags"`
	EventIDFrom          string              `json:"eventid_from"`
	EventIDTill          string              `json:"eventid_till"`
	TimeFrom             string              `json:"time_from"`
	TimeTill             string              `json:"time_till"`
	ProblemTimeFrom      string              `json:"problem_time_from"`
	ProblemTimeTill      string              `json:"problem_time_till"`
	Value                []int               `json:"value"`
	SelectHosts          string              `json:"selectHosts"`
	SelectRelatedObject  string              `json:"selectRelatedObject"`
	SelectAlerts         string              `json:"select_alerts"`
	SelectAcknowledges   string              `json:"select_acknowledges"`
	SelectTags           string              `json:"selectTags"`
	SelectSupressionData string              `json:"selectSupressionData"`
	ZabbixCommun
}

type EventOutput struct {
	EventID        string                `json:"eventid"`
	Source         string                `json:"source"`
	Object         string                `json:"object"`
	ObjectID       string                `json:"objectid"`
	Clock          string                `json:"clock"`
	Value          string                `json:"value"`
	Acknowledged   string                `json:"acknowledged"`
	NS             string                `json:"ns"`
	Name           string                `json:"name"`
	Severity       string                `json:"severity"`
	REventID       string                `json:"r_eventid"`
	CEventID       string                `json:"c_eventid"`
	CorrelationID  string                `json:"correlationid"`
	UserID         string                `json:"userid"`
	CauseEventID   string                `json:"cause_eventid"`
	Opdata         string                `json:"opdata"`
	Acknowledges   []EventAcknowledge    `json:"acknowledges"`
	SupressionData []EventSupressionData `json:"supressiondata"`
	Supressed      string                `json:"supressed"`
	Tags           []map[string]string   `json:"tags"`
}

type EventAcknowledge struct {
	AcknowledgeID string `json:"acknowledgeid"`
	UserID        string `json:"userid"`
	EventID       string `json:"eventid"`
	Clock         string `json:"clock"`
	Message       string `json:"message"`
	Action        string `json:"action"`
	OldSeverity   string `json:"old_severity"`
	NewSeverity   string `json:"new_severity"`
	SupressUntil  string `json:"supress_until"`
	TaskID        string `json:"taskid"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
}

type EventSupressionData struct {
	MaintenanceID string `json:"maintenanceid"`
	SuppressUntil string `json:"suppress_until"`
	UserID        string `json:"userid"`
}
