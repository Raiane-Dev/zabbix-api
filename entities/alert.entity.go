package entities

type AlertGet struct {
	AlertIDs         []string `json:"alertids,omitempty"`
	ActionIDs        []string `json:"actionids,omitempty"`
	EventIDs         []string `json:"eventids,omitempty"`
	GroupIDs         []string `json:"groupids,omitempty"`
	HostIDs          []string `json:"hostids,omitempty"`
	MediaTypeIDs     []string `json:"mediatypeids,omitempty"`
	ObjectIDs        []string `json:"objectids,omitempty"`
	UserIDs          []string `json:"userids,omitempty"`
	EventObject      int      `json:"eventobject,omitempty"`
	EventSource      int      `json:"eventsource,omitempty"`
	TimeFrom         string   `json:"time_from,omitempty"`
	TimeTill         string   `json:"time_till,omitempty"`
	SelectHosts      string   `json:"selectHosts,omitempty"`
	SelectMediaTypes string   `json:"selectMediaTypes,omitempty"`
	SelectUsers      string   `json:"selectUsers,omitempty"`
	ZabbixCommun
}

type AlertOutput struct {
	AlertID       string `json:"alertid,omitempty"`
	ActionID      string `json:"actionid,omitempty"`
	EventID       string `json:"eventid,omitempty"`
	UserID        string `json:"userid,omitempty"`
	Clock         string `json:"clock,omitempty"`
	MediaTypeID   string `json:"mediatypeid,omitempty"`
	SendTo        string `json:"sendto,omitempty"`
	Subject       string `json:"subject,omitempty"`
	Message       string `json:"message,omitempty"`
	Status        string `json:"status,omitempty"`
	Retries       string `json:"retries,omitempty"`
	Error         string `json:"error,omitempty"`
	EscStep       string `json:"escstep,omitempty"`
	AlertType     string `json:"alerttype,omitempty"`
	PEventID      string `json:"p_eventid,omitempty"`
	AcknowledgeID string `json:"acknowledgeid,omitempty"`
}
